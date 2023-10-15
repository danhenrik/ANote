import React from "react";
import { Typography, Select, MenuItem, TextareaAutosize } from "@mui/material";
import { useFormik } from "formik";
import * as yup from "yup";
import {
  InputLabel,
  Button,
  TextField,
} from "../../../../common/FormStyling.styled";
import PropTypes from "prop-types";
import { useModal } from "../../../../store/modal-context";

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

const NoteForm = ({ notes }) => {
  const modal = useModal();
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
      notes.push({
        Id: "6",
        Title: "Note 2",
        Content:
          "Content for Note 2Content for Note 2Content for Note 2Content for Note 2Content for Note 2Content for Note 2Content for Note 2Content for Note 2Content for Note 2Content for Note 2Content for Note 2Content for Note 2Content for Note 2Content for Note 2Content for Note 2",
        LikesCount: 15,
        Likes: ["user2", "user3", "user4"],
        PublishedDate: "2023-10-02T09:45:00Z",
        UpdatedDate: "2023-10-02T14:15:00Z",
        Author: "jane_smith",
        Tags: ["tag2", "tag3", "tag4"],
        CommentCount: 7,
        Commenters: ["user1", "user3", "user5", "user6", "user8"],
      });
      console.log(values);
      modal.closeModal();
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
          {...(formik.touched.description && formik.errors.description
            ? { error: "true" }
            : {})}
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

NoteForm.propTypes = {
  notes: PropTypes.array.isRequired,
};

export default NoteForm;
