import React from "react";
import { useNavigate, useSearchParams } from "react-router-dom";
import {
  InputLabel,
  TextField,
  Button,
} from "../../../common/FormStyling.style";
import { Container, Paper, Typography } from "@mui/material";
import { useFormik } from "formik";
import * as yup from "yup";

const validationSchema = yup.object({
  email: yup
    .string("Insira um email")
    .email("Insira um email válido")
    .required("Insira um email"),
  userId: yup
    .string("Insira um ID de usuário")
    .required("Insira um ID de usuário"),
  password: yup
    .string("Insira uma senha")
    .min(8, "A senha deve ter pelo menos 8 caracteres")
    .required("Insira uma senha"),
  confirmPassword: yup
    .string("Confirme sua senha")
    .oneOf([yup.ref("password"), null], "As senhas devem corresponder")
    .required("Confirme sua senha"),
});

const CompleteSignupForm = () => {
  const [search] = useSearchParams();
  const navigate = useNavigate();
  const formik = useFormik({
    initialValues: {
      email: search.get("email"),
      userId: "",
      password: "",
      confirmPassword: "",
    },
    validationSchema: validationSchema,
    onSubmit: (values) => {
      alert(JSON.stringify(values, null, 2));
      navigate("/signup");
    },
  });

  return (
    <div
      style={{
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        height: "80vh",
      }}
    >
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
              placeholder='Digite um e-mail'
              onBlur={formik.handleBlur}
              error={formik.touched.email && Boolean(formik.errors.email)}
              helperText={formik.touched.email && formik.errors.email}
              autoComplete='email'
            />
            <InputLabel htmlFor='userId'>Id do Usuario</InputLabel>
            <TextField
              fullWidth
              id='userId'
              name='userId'
              value={formik.values.userId}
              onChange={formik.handleChange}
              placeholder='Digite um id de usuario'
              onBlur={formik.handleBlur}
              error={formik.touched.userId && Boolean(formik.errors.userId)}
              helperText={formik.touched.userId && formik.errors.userId}
            />
            <InputLabel htmlFor='password'>Senha</InputLabel>
            <TextField
              fullWidth
              id='password'
              name='password'
              type='password'
              placeholder='Digite uma senha'
              value={formik.values.password}
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
              error={formik.touched.password && Boolean(formik.errors.password)}
              helperText={formik.touched.password && formik.errors.password}
            />
            <InputLabel htmlFor='confirmPassword'>Confirme a Senha</InputLabel>
            <TextField
              fullWidth
              id='confirmPassword'
              name='confirmPassword'
              type='password'
              placeholder='Digite sua senha novamente'
              value={formik.values.confirmPassword}
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
              error={
                formik.touched.confirmPassword &&
                Boolean(formik.errors.confirmPassword)
              }
              helperText={
                formik.touched.confirmPassword && formik.errors.confirmPassword
              }
            />
            <Button variant='contained' fullWidth type='submit'>
              Cadastrar-se
            </Button>
          </form>
        </Paper>
      </Container>
    </div>
  );
};

export default CompleteSignupForm;
