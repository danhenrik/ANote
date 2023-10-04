import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import PropTypes from "prop-types";
import {
  AvatarAuthor,
  AvatarBackground,
  AvatarContainer,
  AvatarUsernames,
  CustomAvatar,
  StyledLink,
  Title,
} from "./NoteCard.styled";

function NoteCard({ note }) {
  var randomColor = require("randomcolor");

  return (
    <Card>
      <CardContent>
        <Title variant='h7' component='div'>
          {note.Title}
        </Title>
        <StyledLink to='/404'>
          <AvatarBackground randomColor={randomColor}>
            <AvatarContainer>
              <CustomAvatar variant='square'>N</CustomAvatar>
            </AvatarContainer>
            <AvatarUsernames>
              <AvatarAuthor color='textSecondary'>{note.Author}</AvatarAuthor>

              <Typography color='textSecondary'>anc</Typography>
            </AvatarUsernames>
          </AvatarBackground>
        </StyledLink>
        <Typography color='textSecondary'>{note.Content}</Typography>
        <Typography color='textSecondary'>
          Published Date: {note.PublishedDate}
        </Typography>
        <Typography variant='body2'>{note.Content}</Typography>
      </CardContent>
    </Card>
  );
}

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

NoteCard.propTypes = {
  note: noteShape.isRequired,
};

export default NoteCard;
