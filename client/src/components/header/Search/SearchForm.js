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
      if (!value.creationDate) return true; // Allow empty input
      // Use a regular expression to match "dd/mm/yyyy" format
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
      const search = {
        ...values,
        tags: tagList,
      };

      alert(JSON.stringify(search, null, 2));
    },
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
              error={formik.touched.title && Boolean(formik.errors.title)}
              helperText={formik.touched.title && formik.errors.title}
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
              error={formik.touched.content && Boolean(formik.errors.content)}
              helperText={formik.touched.content && formik.errors.content}
            />
            <InputLabel htmlFor='user'>Usuário</InputLabel>
            <TextField
              fullWidth
              name='user'
              id='user'
              variant='outlined'
              value={formik.values.user}
              onChange={formik.handleChange}
              error={formik.touched.user && Boolean(formik.errors.user)}
              helperText={formik.touched.user && formik.errors.user}
            />
            <InputLabel htmlFor='community'>Comunidade</InputLabel>
            <TextField
              fullWidth
              name='community'
              id='community'
              variant='outlined'
              value={formik.values.community}
              onChange={formik.handleChange}
              error={
                formik.touched.community && Boolean(formik.errors.community)
              }
              helperText={formik.touched.community && formik.errors.community}
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
