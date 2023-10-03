import React from "react";
import { useNavigate } from "react-router-dom";
import {
  InputLabel,
  TextField,
  Button,
} from "../../../common/FormStyling.style";
import { Container, Paper, Typography } from "@mui/material";
import { useFormik } from "formik";
import * as yup from "yup";

const validationSchema = yup.object({
  password: yup
    .string("Insira uma senha")
    .min(8, "A senha deve ter pelo menos 8 caracteres")
    .required("Insira uma senha"),
  confirmPassword: yup
    .string("Confirme sua senha")
    .oneOf([yup.ref("password"), null], "As senhas devem corresponder")
    .required("Confirme sua senha"),
});

const CompleteRecovery = () => {
  const navigate = useNavigate();
  const formik = useFormik({
    initialValues: {
      password: "",
      confirmPassword: "",
    },
    validationSchema: validationSchema,
    onSubmit: (values) => {
      alert(JSON.stringify(values, null, 2));
      navigate("/");
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
            Recuperar Senha
          </Typography>
          <form onSubmit={formik.handleSubmit} style={{ width: "100%" }}>
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
              Recuperar Senha
            </Button>
          </form>
        </Paper>
      </Container>
    </div>
  );
};

export default CompleteRecovery;
