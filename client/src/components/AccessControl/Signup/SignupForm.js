import React from "react";
import {
  Container,
  Divider,
  IconButton,
  Paper,
  Typography,
} from "@mui/material";
import { useFormik } from "formik";
import * as yup from "yup";
import {
  InputLabel,
  Button,
  TextField,
} from "../../../common/FormStyling.styled";
import SocialMediaAuth from "../SocialMediaAuth";
import { PropTypes } from "prop-types";
import { createSearchParams, useNavigate } from "react-router-dom";
import CloseIcon from "@mui/icons-material/Close";

const validationSchema = yup.object({
  email: yup
    .string("Insira seu email")
    .email("Insira um email válido")
    .required("Insira seu email"),
});

const SignupForm = ({ closeModal }) => {
  const navigate = useNavigate();
  const formik = useFormik({
    initialValues: {
      email: "",
    },
    validationSchema: validationSchema,
    onSubmit: (values) => {
      closeModal();
      alert(JSON.stringify(values, null, 2));
      navigate({
        pathname: "signup",
        search: createSearchParams({
          email: values.email,
        }).toString(),
      });
    },
  });

  return (
    <Container maxWidth='xs'>
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
        <Typography variant='h6' style={{ fontWeight: "bold" }}>
          Cadastre-se
        </Typography>
        <form onSubmit={formik.handleSubmit} style={{ width: "100%" }}>
          <InputLabel htmlFor='email'>E-Mail</InputLabel>
          <TextField
            fullWidth
            id='email'
            name='email'
            value={formik.values.email}
            onChange={formik.handleChange}
            placeholder='Digite seu e-mail'
            onBlur={formik.handleBlur}
            error={formik.touched.email && Boolean(formik.errors.email)}
            helperText={formik.touched.email && formik.errors.email}
          />
          <Button variant='contained' fullWidth type='submit'>
            Cadastrar-se
          </Button>
          <Divider sx={{ mt: "30px", mb: "30px" }}>Ou</Divider>
          <SocialMediaAuth authType='Cadastro'></SocialMediaAuth>
        </form>
      </Paper>
    </Container>
  );
};

SignupForm.propTypes = {
  closeModal: PropTypes.func.isRequired,
};

export default SignupForm;
