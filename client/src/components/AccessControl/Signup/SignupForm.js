import React from "react";
import { Divider, Typography } from "@mui/material";
import { useFormik } from "formik";
import * as yup from "yup";
import {
  InputLabel,
  Button,
  TextField,
} from "../../../common/FormStyling.styled";
import SocialMediaAuth from "../SocialMediaAuth";
import { createSearchParams, useNavigate } from "react-router-dom";
import { useModal } from "../../../store/modal-context";

const validationSchema = yup.object({
  email: yup
    .string("Insira seu email")
    .email("Insira um email válido")
    .required("Insira seu email"),
});

const SignupForm = () => {
  const navigate = useNavigate();
  const modal = useModal();
  const formik = useFormik({
    initialValues: {
      email: "",
    },
    validationSchema: validationSchema,
    onSubmit: (values) => {
      modal.closeModal();
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
    <>
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
    </>
  );
};

export default SignupForm;
