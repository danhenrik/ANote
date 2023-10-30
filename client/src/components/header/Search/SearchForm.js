import React, { useEffect, useState } from "react";
import { FormikProvider, useFormik } from "formik";
import { Container, IconButton, Paper, Typography } from "@mui/material";
import * as yup from "yup";
import {
  InputLabel,
  Button,
  TextField,
  TextArea,
} from "../../../common/FormStyling.styled";
import CloseIcon from "@mui/icons-material/Close";
import { PropTypes } from "prop-types";
import ReactInputMask from "react-input-mask";
import InputAutocomplete from "../../../common/InputAutoComplete";
import Tags from "../../Timeline/Tags/TagsList";
import useTags from "../../../api/useTags";
import listHandler from "./SearchListHandler";
import { createSearchParams, useNavigate } from "react-router-dom";

const dynamicValidation = (tagList) => (value) => {
  if (
    value.title ||
    value.content ||
    value.user ||
    tagList.length > 0 ||
    value.community ||
    value.creationDate
  ) {
    return true;
  }

  return new yup.ValidationError(
    "Insira pelo menos um campo",
    "null",
    "atleastone",
    "required"
  );
};

export const validationSchema = (tagList) => {
  return yup
    .object()
    .shape({
      title: yup.string(),
      content: yup.string(),
      user: yup.string(),
      community: yup.string(),
      creationDate: yup.string(),
    })
    .test("back", "Invalid input", dynamicValidation(tagList))
    .test("valid-date", "Invalid date format", (value) => {
      if (!value.creationDate) return true;
      const dateRegex = /^(\d{2})\/(\d{2})\/(\d{4})$/;
      if (!dateRegex.test(value.creationDate))
        return new yup.ValidationError(
          "A data deve estar no formato dd/mm/aaaa",
          "null",
          "date",
          "required"
        );
      const [, day, month, year] = dateRegex.exec(value.creationDate);
      const parsedDate = new Date(`${year}-${month}-${day}`);
      if (parsedDate.toString() == "Invalid Date")
        return new yup.ValidationError(
          "Insira uma data válida no formato dd/mm/aaaa",
          "null",
          "date",
          "required"
        );
    });
};

const SearchForm = ({ closeModal }) => {
  const [tagList, setTagList] = useState([]);
  const [tags, setTags] = useState([]);
  const tagsApi = useTags();
  const navigate = useNavigate();

  useEffect(() => {
    const fetchTags = async () => {
      let fetchedTags = await tagsApi.fetchTags();
      fetchedTags = fetchedTags.map((item) => item.Tags);
      setTags(fetchedTags);
    };
    fetchTags();
  }, []);

  const formik = useFormik({
    initialValues: {
      title: "",
      content: "",
      user: "",
      community: "",
      creationDate: "",
    },
    validationSchema: validationSchema(tagList),
    onSubmit: (values) => {
      const queryParams = [];
      for (const key in values) {
        if (key !== "tags") {
          if (values[key] != "")
            queryParams.push(`${key}=${encodeURIComponent(values[key])}`);
        }
      }

      for (const tag of tagList) {
        queryParams.push(`tags=${encodeURIComponent(tag)}`);
      }

      const searchParams = queryParams.join("&");

      navigate(`/timeline?search=true&${searchParams}`);
      closeModal();
    },
    validateOnChange: false,
    validateOnBlur: false,
  });

  const { addToList, removeFromList } = listHandler(setTagList);

  return (
    <FormikProvider value={formik}>
      <Container maxWidth='sm'>
        <Paper
          elevation={3}
          sx={{
            padding: 2,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <IconButton onClick={closeModal} sx={{ alignSelf: "flex-end" }}>
            <CloseIcon />
          </IconButton>
          <Typography variant='h6'>Pesquisar</Typography>
          <form onSubmit={formik.handleSubmit} style={{ width: "100%" }}>
            <InputLabel htmlFor='title'>Título</InputLabel>
            <TextField
              fullWidth
              name='title'
              id='title'
              value={formik.values.title}
              onChange={formik.handleChange}
            />
            <InputLabel htmlFor='content'>Conteúdo</InputLabel>
            <TextArea
              fullWidth
              name='content'
              id='content'
              variant='outlined'
              value={formik.values.content}
              onChange={formik.handleChange}
              multiline
              rows={4}
            />
            <InputLabel htmlFor='user'>Usuário</InputLabel>
            <TextField
              fullWidth
              name='user'
              id='user'
              variant='outlined'
              value={formik.values.user}
              onChange={formik.handleChange}
            />
            <InputLabel htmlFor='community'>Comunidade</InputLabel>
            <TextField
              fullWidth
              name='community'
              id='community'
              variant='outlined'
              value={formik.values.community}
              onChange={formik.handleChange}
            />
            <InputLabel htmlFor='tag'>Tag</InputLabel>
            <InputAutocomplete
              addToList={addToList}
              name='tag'
              id='tag'
              options={tags}
              list={tagList}
            />
            <Tags
              tags={tagList}
              hasLink={false}
              hasDelete={true}
              deletionHandler={removeFromList}
            ></Tags>
            <InputLabel htmlFor='creationDate'>Data de Criação</InputLabel>
            <ReactInputMask
              style={{
                border: "1px solid #ccc",
                borderRadius: "6px",
                padding: "8px",
                height: "40px",
                width: "100%",
              }}
              name='creationDate'
              id='creationDate'
              value={formik.values.creationDate}
              onChange={formik.handleChange}
              mask='99/99/9999'
            >
              {(inputProps) => <input {...inputProps} />}
            </ReactInputMask>
            {formik.errors.date ? (
              <Typography color='error'>{formik.errors.date}</Typography>
            ) : null}
            {formik.errors.atleastone ? (
              <Typography color='error'>{formik.errors.atleastone}</Typography>
            ) : null}
            <Button
              type='submit'
              variant='contained'
              color='primary'
              sx={{ width: "100%" }}
            >
              Pesquisar
            </Button>
          </form>
        </Paper>
      </Container>
    </FormikProvider>
  );
};

SearchForm.propTypes = {
  closeModal: PropTypes.func.isRequired,
};

export default SearchForm;
