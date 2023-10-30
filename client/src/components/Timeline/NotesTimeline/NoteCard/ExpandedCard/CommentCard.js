import React, { useState } from "react";
import Typography from "@mui/material/Typography";
import DeleteIcon from "@mui/icons-material/Delete";
import PropTypes from "prop-types";
import {
  AvatarAuthor,
  AvatarBackground,
  AvatarContainer,
  AvatarUsernames,
  CustomAvatar,
} from "../NoteCard.styled";
import axios from "axios";
import { Card, Container } from "@mui/material";
import { useAuth } from "../../../../../store/auth-context";
import formatDate from "../../../../../util/formatDate";

const CommentCard = ({ comment, numberCommentsHandler }) => {
  const userAuth = useAuth();
  const [renderComment, setRenderComment] = useState(true);
  const formatedDate = formatDate(comment.CreatedAt);

  const deleteComment = () => {
    try {
      axios.delete("/comments/" + comment.Id);
      numberCommentsHandler((prev) => prev - 1);
      setRenderComment(false);
    } catch (error) {
      console.log("Comment delete failed: ", error);
    }
  };

  return (
    <>
      {renderComment && (
        <Card
          sx={{
            minWidth: "100%",
            backgroundColor: "lightgray",
          }}
        >
          <AvatarBackground sx={{ height: "40px" }}>
            <AvatarContainer style={{ height: "40px" }}>
              <CustomAvatar
                style={{ height: "40px", width: "40px" }}
                variant='square'
              >
                N
              </CustomAvatar>
            </AvatarContainer>
            <AvatarUsernames>
              <AvatarAuthor>{comment.Author}</AvatarAuthor>
            </AvatarUsernames>
            {comment.Author === userAuth.username && (
              <DeleteIcon
                onClick={deleteComment}
                style={{
                  cursor: "pointer",
                  color: "red",
                  display: "block",
                  marginRight: "15px",
                  marginLeft: "auto",
                }}
              />
            )}
          </AvatarBackground>
          <Container sx={{ marginTop: "20px", marginBottom: "20px" }}>
            <Typography color='textSecondary'>{comment.Content}</Typography>
          </Container>
          <Typography sx={{ textAlign: "right" }} color='textSecondary'>
            {formatedDate.day} às {formatedDate.hour}
          </Typography>
        </Card>
      )}
    </>
  );
};

const commentShape = PropTypes.shape({
  Id: PropTypes.string.isRequired,
  Author: PropTypes.string.isRequired,
  Content: PropTypes.string.isRequired,
  CreatedAt: PropTypes.string.isRequired,
});

CommentCard.propTypes = {
  comment: commentShape.isRequired,
  numberCommentsHandler: PropTypes.func,
};

export default CommentCard;
