import React from "react";
import { Typography, Select, MenuItem, TextareaAutosize } from "@mui/material";
import { useFormik } from "formik";
import * as yup from "yup";
import {
  InputLabel,
  Button,
  TextField,
} from "../../../../common/FormStyling.styled";
const validationSchema = yup.object({
  title: yup.string("Insira o título").required("Título é obrigatório"),
  description: yup
    .string("Insira a descrição")
    .required("Descrição é obrigatória"),
  tags: yup.string("Insira as tags"),
  privacy: yup
    .string("Selecione a privacidade")
    .required("Privacidade é obrigatória"),
});

const NoteForm = () => {
  //const navigate = useNavigate();
  const formik = useFormik({
    initialValues: {
      title: "",
      description: "",
      tags: "",
      privacy: "public",
    },
    validationSchema: validationSchema,
    onSubmit: (values) => {
      console.log(values);
    },
  });

  return (
    <>
      <Typography variant='h6' style={{ fontWeight: "bold" }}>
        Criar uma Nota
      </Typography>
      <form onSubmit={formik.handleSubmit} style={{ width: "100%" }}>
        <InputLabel htmlFor='title'>Título</InputLabel>
        <TextField
          fullWidth
          id='title'
          name='title'
          placeholder='Insira o título'
          value={formik.values.title}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          error={formik.touched.title && Boolean(formik.errors.title)}
          helperText={formik.touched.title && formik.errors.title}
        />
        <InputLabel htmlFor='description'>Descrição</InputLabel>
        <TextareaAutosize
          id='description'
          name='description'
          placeholder='Insira a descrição'
          value={formik.values.description}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          error={
            formik.touched.description && Boolean(formik.errors.description)
          }
          minRows={4}
          style={{
            width: "100%",
            border: "1px solid #ced4da",
            borderRadius: "4px",
            padding: "10px",
            fontSize: "16px",
          }}
        />
        <InputLabel htmlFor='tags'>Tags</InputLabel>
        <TextField
          fullWidth
          id='tags'
          name='tags'
          placeholder='Insira as tags'
          value={formik.values.tags}
          onChange={formik.handleChange}
        />
        <InputLabel htmlFor='privacy'>Privacidade</InputLabel>
        <Select
          fullWidth
          id='privacy'
          name='privacy'
          value={formik.values.privacy}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          error={formik.touched.privacy && Boolean(formik.errors.privacy)}
        >
          <MenuItem value='public'>Público</MenuItem>
          <MenuItem value='private'>Privado</MenuItem>
        </Select>
        <Button variant='contained' fullWidth type='submit'>
          Criar Nota
        </Button>
      </form>
    </>
  );
};

export default NoteForm;
