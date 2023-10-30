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
  code: yup
    .string()
    .matches(/^[0-9]+$/, "O código deve ser composto apenas por números")
    .min(4, "O código deve ter 4 dígitos")
    .max(4, "O código deve ter 4 dígitos"),
});

const VerificationCode = () => {
  const navigate = useNavigate();
  const formik = useFormik({
    initialValues: {
      code: "",
    },
    validationSchema: validationSchema,
    onSubmit: (values) => {
      navigate("/passwordrecovery/completerecovery");
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
            <InputLabel htmlFor='code'>Insira o Código</InputLabel>
            <TextField
              fullWidth
              id='code'
              name='code'
              value={formik.values.code}
              onChange={formik.handleChange}
              placeholder='Digite o código de recuperação'
              onBlur={formik.handleBlur}
              error={formik.touched.code && Boolean(formik.errors.code)}
              helperText={formik.touched.code && formik.errors.code}
              autoComplete='code'
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

export default VerificationCode;
