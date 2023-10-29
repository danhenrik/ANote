import React from "react";
import { Typography } from "@mui/material";
import { useFormik } from "formik";
import * as yup from "yup";
import {
  InputLabel,
  Button,
  TextField,
} from "../../../../common/FormStyling.styled";
import PropTypes from "prop-types";
import { useModal } from "../../../../store/modal-context";
import useCommunities from "../../../../api/useCommunities";

const validationSchema = yup.object({
  name: yup.string("Insira o nome").required("Nome é obrigatório"),
});

const CommunityForm = ({ communities, setCommunitiesHandler }) => {
  const communitiesApi = useCommunities();
  const modal = useModal();
  const formik = useFormik({
    initialValues: {
      name: "",
    },
    validationSchema: validationSchema,
    onSubmit: (values) => {
      const community = {
        name: values.name,
      };
      const createCommunity = async () => {
        const createdCommunity =
          await communitiesApi.createCommunity(community);
        if (createdCommunity) {
          communities.push(community);
          setCommunitiesHandler(communities);
        }
      };

      createCommunity();

      modal.closeModal();
    },
  });

  return (
    <>
      <Typography variant='h6' style={{ fontWeight: "bold" }}>
        Criar uma Comunidade
      </Typography>
      <form onSubmit={formik.handleSubmit} style={{ width: "100%" }}>
        <InputLabel htmlFor='name'>Nome</InputLabel>
        <TextField
          fullWidth
          id='name'
          name='name'
          placeholder='Insira o nome'
          value={formik.values.name}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          error={formik.touched.name && Boolean(formik.errors.name)}
          helperText={formik.touched.name && formik.errors.name}
        />
        <Button variant='contained' fullWidth type='submit'>
          Criar Comunidade
        </Button>
      </form>
    </>
  );
};

CommunityForm.propTypes = {
  communities: PropTypes.array.isRequired,
  setCommunitiesHandler: PropTypes.func.isRequired,
};

export default CommunityForm;
