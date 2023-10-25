import React from "react";
import { Typography, IconButton } from "@mui/material";
import { useFormik } from "formik";
import * as yup from "yup";
import {
  InputLabel,
  Button,
  TextField,
} from "../../../../common/FormStyling.styled";
import PropTypes from "prop-types";
import { useModal } from "../../../../store/modal-context";
import { CustomTextArea } from "../../TimelineForms.styled";
import AddIcon from "@mui/icons-material/Add";
import useNotes from "../../../../api/useNotes";

const validationSchema = yup.object({
  title: yup.string("Insira o título").required("Título é obrigatório"),
  description: yup
    .string("Insira a descrição")
    .required("Descrição é obrigatória"),
  tags: yup.string("Insira as tags"),
});

const NoteForm = ({ notes }) => {
  const notesApi = useNotes();
  const modal = useModal();
  //const navigate = useNavigate();
  const formik = useFormik({
    initialValues: {
      title: "",
      description: "",
      tags: "",
    },
    validationSchema: validationSchema,
    onSubmit: (values) => {
      const note = {
        title: values.title,
        description: values.description,
        tags: [values.tags],
      };
      const postNotes = async () => {
        const fetchedNotes = await notesApi.createNote(note);
        if (fetchedNotes) {
          notes.push(fetchedNotes);
        }
      };
      postNotes();
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
        <CustomTextArea
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
        />
        <InputLabel htmlFor='tags'>Tags</InputLabel>
        <div style={{ display: "flex", alignItems: "center" }}>
          <TextField
            fullWidth
            id='tags'
            name='tags'
            placeholder='Insira as tags'
            value={formik.values.tags}
            onChange={formik.handleChange}
          />
          <IconButton style={{ marginLeft: "5px" }}>
            <AddIcon />
          </IconButton>
        </div>
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
