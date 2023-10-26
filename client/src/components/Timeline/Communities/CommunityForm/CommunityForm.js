import React from "react";
import { Typography, MenuItem, IconButton } from "@mui/material";
import { useFormik } from "formik";
import * as yup from "yup";
import {
  InputLabel,
  Button,
  TextField,
} from "../../../../common/FormStyling.styled";
import PropTypes from "prop-types";
import { useModal } from "../../../../store/modal-context";
import { CustomSelect } from "../../TimelineForms.styled";
import AddIcon from "@mui/icons-material/Add";
import useCommunities from "../../../../api/useCommunities";

const validationSchema = yup.object({
  name: yup.string("Insira o nome").required("Nome é obrigatório"),
  tags: yup.string("Insira as tags"),
});

const CommunityForm = ({ communities }) => {
  const communitiesApi = useCommunities();
  const modal = useModal();
  const formik = useFormik({
    initialValues: {
      name: "",
      tags: "",
    },
    validationSchema: validationSchema,
    onSubmit: (values) => {
      const community = {
        name: values.name,
        tags: values.tags,
      };
      const createCommunity = async () => {
        const createdCommunity =
          await communitiesApi.createCommunity(community);
        if (createdCommunity) {
          console.log("criado");
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
        <InputLabel htmlFor='tags'>Tags</InputLabel>
        <div style={{ display: "flex", alignItems: "center" }}>
          <TextField
            fullWidth
            id='tags'
            name='tags'
            placeholder='Insira as tags'
            value={formik.values.tags}
            onChange={formik.handleChange}
          />
          <IconButton style={{ marginLeft: "5px" }}>
            <AddIcon />
          </IconButton>
        </div>
        <Button variant='contained' fullWidth type='submit'>
          Criar Comunidade
        </Button>
      </form>
    </>
  );
};

CommunityForm.propTypes = {
  communities: PropTypes.array.isRequired,
};

export default CommunityForm;
