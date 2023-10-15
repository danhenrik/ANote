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
import { Link, useNavigate } from "react-router-dom";
import { PropTypes } from "prop-types";
import { useAuth } from "../../../store/auth-context";
import { useModal } from "../../../store/modal-context";

const validationSchema = yup.object({
  email: yup
    .string("Insira seu email")
    .email("Insira um e-mail válido")
    .required("Insira seu e-mail"),
  password: yup
    .string("Insira sua senha")
    .min(8, "A senha deve conter pelo menos 8 caracteres, com uma letra")
    .required("Insira sua senha"),
});

const LoginForm = () => {
  const auth = useAuth();
  const navigate = useNavigate();
  const modal = useModal();
  const formik = useFormik({
    initialValues: {
      email: "",
      password: "",
    },
    validationSchema: validationSchema,
    onSubmit: (values) => {
      auth.login(values, "EMAIL");
      modal.closeModal();
      navigate("/");
    },
  });

  return (
    <>
      <Typography variant='h6' style={{ fontWeight: "bold" }}>
        Faça Login
      </Typography>
      <form onSubmit={formik.handleSubmit} style={{ width: "100%" }}>
        <InputLabel htmlFor='email'>E-Mail ou Usuário</InputLabel>
        <TextField
          fullWidth
          id='email'
          name='email'
          placeholder='Digite seu email'
          value={formik.values.email}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          error={formik.touched.email && Boolean(formik.errors.email)}
          helperText={formik.touched.email && formik.errors.email}
        />
        <InputLabel htmlFor='password'>Senha</InputLabel>
        <TextField
          fullWidth
          id='password'
          name='password'
          type='password'
          placeholder='Digite sua senha'
          value={formik.values.password}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          error={formik.touched.password && Boolean(formik.errors.password)}
          helperText={formik.touched.password && formik.errors.password}
        />
        <Button variant='contained' fullWidth type='submit'>
          Fazer Login
        </Button>
        <div style={{ marginTop: "10px" }}>
          <Link
            style={{ color: "blue" }}
            to='/passwordrecovery'
            onClick={modal.closeModal}
          >
            Esqueceu sua senha?
          </Link>
        </div>
        <Divider sx={{ mt: "30px", mb: "15px" }}>Ou</Divider>
        <SocialMediaAuth authType='Login'></SocialMediaAuth>
      </form>
    </>
  );
};

export default LoginForm;
