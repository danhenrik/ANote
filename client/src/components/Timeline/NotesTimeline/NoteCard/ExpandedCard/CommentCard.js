import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import DeleteIcon from "@mui/icons-material/Delete";
import PropTypes from "prop-types";
import {
  AvatarAuthor,
  AvatarBackground,
  AvatarContainer,
  AvatarUsernames,
  ContentContainer,
  CustomAvatar,
  StyledLink,
} from "../NoteCard.styled";

import { Card } from "@mui/material";

const CommentCard = ({ comment }) => {
  return (
    <Card sx={{ minWidth: "100%", backgroundColor: "#c0c0c0" }}>
      <CardContent>
        <AvatarBackground>
          <AvatarContainer>
            <CustomAvatar variant='square'>N</CustomAvatar>
          </AvatarContainer>
          <StyledLink to='/404'>
            <AvatarUsernames>
              <AvatarAuthor>{comment.Author}</AvatarAuthor>
            </AvatarUsernames>
          </StyledLink>
          <DeleteIcon
            style={{
              cursor: "pointer",
              color: "red",
              display: "block",
              marginRight: "15px",
              marginLeft: "auto",
            }}
          />
        </AvatarBackground>
        <ContentContainer sx={{ marginTop: "10px" }}>
          <Typography color='textSecondary'>{comment.Content}</Typography>
        </ContentContainer>
        <Typography color='textSecondary'>{comment.CreatedAt}</Typography>
      </CardContent>
    </Card>
  );
};

const commentShape = PropTypes.shape({
  Author: PropTypes.string.isRequired,
  Content: PropTypes.string.isRequired,
  CreatedAt: PropTypes.string.isRequired,
});

CommentCard.propTypes = {
  comment: commentShape.isRequired,
};

export default CommentCard;
