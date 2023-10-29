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
import { useAuth } from "../../../store/auth-context";
import { useModal } from "../../../store/modal-context";

const validationSchema = yup.object({
  login: yup
    .mixed()
    .test(
      "is-email-or-username",
      "Insira um e-mail válido ou nome de usuário",
      (value) => {
        if (!value) return false; // No value provided, fail the test
        return isValidEmail(value) || isValidUsername(value);
      }
    )
    .required("Insira seu e-mail ou nome de usuário"),
});

function isValidEmail(value) {
  return /\S+@\S+\.\S+/.test(value);
}

function isValidUsername(value) {
  return /^[a-zA-Z0-9_]+$/.test(value);
}

const LoginForm = () => {
  const auth = useAuth();
  const navigate = useNavigate();
  const modal = useModal();
  const formik = useFormik({
    initialValues: {
      login: "",
      password: "",
    },
    validationSchema: validationSchema,
    onSubmit: (values) => {
      const user = {
        login: values.login,
        password: values.password,
      };
      auth.login(user, "EMAIL");
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
        <InputLabel htmlFor='login'>E-Mail ou Usuário</InputLabel>
        <TextField
          fullWidth
          id='login'
          name='login'
          placeholder='Digite seu login'
          value={formik.values.login}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          error={formik.touched.login && Boolean(formik.errors.login)}
          helperText={formik.touched.login && formik.errors.login}
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
          <Link
            style={{ color: "blue", float: "right" }}
            to='/signup'
            onClick={modal.closeModal}
          >
            Cadastre-se
          </Link>
        </div>
        <Divider sx={{ mt: "30px", mb: "15px" }}>Ou</Divider>
        <SocialMediaAuth authType='Login'></SocialMediaAuth>
      </form>
    </>
  );
};

export default LoginForm;
