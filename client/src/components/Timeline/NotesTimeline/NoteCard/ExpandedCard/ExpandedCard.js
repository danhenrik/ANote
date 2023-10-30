import Typography from "@mui/material/Typography";
import PropTypes from "prop-types";
import {
  AvatarAuthor,
  AvatarBackground,
  AvatarContainer,
  AvatarUsernames,
  CommentContainer,
  ContentContainer,
  CustomAvatar,
  Title,
} from "../NoteCard.styled";
import Tags from "../../../Tags/TagsList";
import { CreateButton } from "../../../TimelineList.styled";
import CommentCard from "./CommentCard";
import { Box, Grid } from "@mui/material";
import { useEffect, useState } from "react";
import { useFormik } from "formik";
import * as yup from "yup";
import axios from "axios";
import { useAuth } from "../../../../../store/auth-context";
import formatDate from "../../../../../util/formatDate";
import { CustomTextArea } from "../../../TimelineForms.styled";

const validationSchema = yup.object({
  comment: yup.string("Insira um comentário").required("Insira um comentário"),
});

const ExpandedCard = ({ note, randomColorElement, numberCommentsHandler }) => {
  let [comments, setComments] = useState([]);
  const userAuth = useAuth();
  const formatedDate = formatDate(note.PublishedDate);

  const updateNoteComments = async () => {
    try {
      const noteComments = await axios.get("/comments/" + note.Id);
      setComments(noteComments.data.data);
      numberCommentsHandler(noteComments.data.data.length);
    } catch (error) {
      console.log("Comments retrieving failed: ", error);
    }
  };

  useEffect(() => {
    updateNoteComments();
  }, []);

  const formik = useFormik({
    initialValues: {
      comment: "",
    },
    validationSchema: validationSchema,
    onSubmit: async (values, { resetForm }) => {
      try {
        const commentData = {
          user_id: userAuth.username,
          note_id: note.Id,
          content: values.comment,
        };

        resetForm({ values: "" });
        await axios.post("/comments", commentData);
        updateNoteComments();
      } catch (error) {
        console.error("Comment failed: ", error);
      }
    },
  });

  return (
    <div style={{ width: "100%" }}>
      <Title variant='h7' component='div' textAlign='center'>
        {note.Title}
      </Title>
      <AvatarBackground randomColor={randomColorElement}>
        <AvatarContainer>
          <CustomAvatar variant='square'>N</CustomAvatar>
        </AvatarContainer>
        <AvatarUsernames>
          <AvatarAuthor>{note.Author}</AvatarAuthor>
        </AvatarUsernames>
      </AvatarBackground>
      <ContentContainer sx={{ marginTop: "10px" }}>
        <Typography color='textSecondary'>{note.Content}</Typography>
      </ContentContainer>

      <Typography color='textSecondary' textAlign='center'>
        {formatedDate.day} às {formatedDate.hour}
      </Typography>
      <Typography
        variant='h7'
        color='textPrimary'
        sx={{
          textAlign: "left",
          width: "100%",
          textDecoration: "underline",
          fontWeight: "bold",
        }}
      >
        Comentários
      </Typography>
      {userAuth.isAuthenticated && (
        <form onSubmit={formik.handleSubmit}>
          <CustomTextArea
            label='Comentário'
            variant='standard'
            sx={{
              display: "flex",
              margin: "auto",
            }}
            id='comment'
            name='comment'
            value={formik.values.comment}
            onChange={formik.handleChange}
            placeholder='Digite um comentário'
            minRows={2}
          />
          <CreateButton
            sx={{
              marginTop: "10px",
              display: "block",
              marginLeft: "auto",
            }}
            type='submit'
          >
            Comentar
          </CreateButton>
        </form>
      )}
      <CommentContainer sx={{ marginTop: "15px" }}>
        {comments ? (
          comments.map((comment) => (
            <Box key={comment.Id} sx={{ marginBottom: "15px" }}>
              <CommentCard
                numberCommentsHandler={numberCommentsHandler}
                comment={comment}
              />
            </Box>
          ))
        ) : (
          <Typography variant='h7' color='textPrimary'>
            Nenhum comentário nessa nota
          </Typography>
        )}
      </CommentContainer>
    </div>
  );
};

const noteShape = PropTypes.shape({
  Id: PropTypes.string.isRequired,
  Title: PropTypes.string.isRequired,
  Content: PropTypes.string.isRequired,
  LikesCount: PropTypes.number.isRequired,
  Likes: PropTypes.arrayOf(PropTypes.string).isRequired,
  PublishedDate: PropTypes.string.isRequired,
  UpdatedDate: PropTypes.string.isRequired,
  Author: PropTypes.string.isRequired,
  Tags: PropTypes.arrayOf(PropTypes.string).isRequired,
  CommentCount: PropTypes.number.isRequired,
  Commenters: PropTypes.arrayOf(PropTypes.string).isRequired,
});

ExpandedCard.propTypes = {
  note: noteShape.isRequired,
  randomColorElement: PropTypes.string.isRequired,
  numberComments: PropTypes.number,
  numberCommentsHandler: PropTypes.func,
};

export default ExpandedCard;
