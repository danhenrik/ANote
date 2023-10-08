import React from "react";
import { useNavigate } from "react-router-dom";
import {
  InputLabel,
  TextField,
  Button,
} from "../../../common/FormStyling.styled";
import { Container, Paper, Typography } from "@mui/material";
import { useFormik } from "formik";
import * as yup from "yup";

const validationSchema = yup.object({
  email: yup
    .string("Insira um email")
    .email("Insira um email válido")
    .required("Insira um email"),
});

const PasswordRecovery = () => {
  const navigate = useNavigate();
  const formik = useFormik({
    initialValues: {
      email: "",
    },
    validationSchema: validationSchema,
    onSubmit: (values) => {
      alert(JSON.stringify(values, null, 2));
      navigate("verificationcode", { relative: "path" });
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
            <Button variant='contained' fullWidth type='submit'>
              Recuperar Senha
            </Button>
          </form>
        </Paper>
      </Container>
    </div>
  );
};

export default PasswordRecovery;
