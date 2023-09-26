import React from "react";
import { Container, Divider, Paper, Typography } from "@mui/material";
import { useFormik } from "formik";
import * as yup from "yup";
import {
  InputLabel,
  Button,
  TextField,
} from "../../../common/FormStyling.style";
import SocialMediaAuth from "../login/SocialMediaAuth";
import { PropTypes } from "prop-types";
import { useNavigate } from "react-router-dom";

const validationSchema = yup.object({
  email: yup
    .string("Insira seu email")
    .email("Insira um email válido")
    .required("Insira seu email"),
});

const googleHandler = () => {
  alert("hello");
};
const facebookHandler = () => {
  alert("hello");
};

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
      navigate("/signup");
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
          <SocialMediaAuth
            authType='Cadastro'
            googleHandler={googleHandler}
            facebookHandler={facebookHandler}
          ></SocialMediaAuth>
        </form>
      </Paper>
    </Container>
  );
};
SignupForm.propTypes = {
  closeModal: PropTypes.func.isRequired,
};
export default SignupForm;
