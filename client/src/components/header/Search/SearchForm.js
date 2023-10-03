import React from "react";
import { useFormik } from "formik";
import { Container, IconButton, Paper, Typography } from "@mui/material";
import * as yup from "yup";
import {
  InputLabel,
  Button,
  TextField,
  TextArea,
} from "../../../common/FormStyling.style";
import CloseIcon from "@mui/icons-material/Close";
import { PropTypes } from "prop-types";

export const validationSchema = yup
  .object()
  .shape({
    title: yup.string(),
    content: yup.string(),
    user: yup.string(),
    tag: yup.string(),
    community: yup.string(),
    creationDate: yup.string(),
  })
  .test("back", function _(value) {
    const a = !!(
      value.title ||
      value.content ||
      value.user ||
      value.tag ||
      value.community
    );
    if (!a) {
      let errorMsg = "";
      if (!value.creationDate) {
        errorMsg = "Preencha pelo menos um campo";
      } else {
        errorMsg = "Preencha pelo menos um campo (além da data)";
      }
      return new yup.ValidationError(
        errorMsg,
        "null",
        "atleastone",
        "required"
      );
    }
    return true;
  });

const SearchForm = ({ closeModal }) => {
  const formik = useFormik({
    initialValues: {
      title: "",
      content: "",
      user: "",
      tag: "",
      community: "",
      creationDate: "",
    },
    validationSchema: validationSchema,
    onSubmit: (values) => {
      alert(JSON.stringify(values, null, 2));
    },
  });

  return (
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
          <InputLabel htmlFor='tag'>Tag</InputLabel>
          <TextField
            fullWidth
            name='tag'
            id='tag'
            variant='outlined'
            value={formik.values.tag}
            onChange={formik.handleChange}
            error={formik.touched.tag && Boolean(formik.errors.tag)}
            helperText={formik.touched.tag && formik.errors.tag}
          />
          <InputLabel htmlFor='community'>Comunidade</InputLabel>
          <TextField
            fullWidth
            name='community'
            id='community'
            variant='outlined'
            value={formik.values.community}
            onChange={formik.handleChange}
            error={formik.touched.community && Boolean(formik.errors.community)}
            helperText={formik.touched.community && formik.errors.community}
          />
          <InputLabel htmlFor='creationDate'>Data de Criação</InputLabel>
          <TextField
            fullWidth
            name='creationDate'
            id='creationDate'
            variant='outlined'
            value={formik.values.creationDate}
            onChange={formik.handleChange}
            error={
              formik.touched.creationDate && Boolean(formik.errors.creationDate)
            }
            helperText={
              formik.touched.creationDate && formik.errors.creationDate
            }
          />
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
  );
};

SearchForm.propTypes = {
  closeModal: PropTypes.func.isRequired,
};

export default SearchForm;
