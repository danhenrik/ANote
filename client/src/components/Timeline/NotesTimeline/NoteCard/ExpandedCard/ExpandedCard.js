import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import TextField from "@mui/material/TextField";
import PropTypes from "prop-types";
import {
  AvatarAuthor,
  AvatarBackground,
  AvatarContainer,
  AvatarUsernames,
  ContentContainer,
  CustomAvatar,
  NotesCardContainer,
  StyledLink,
  Title,
} from "../NoteCard.styled";
import Tags from "../../../Tags/TagsList";
import { useAuth } from "../../../../../store/auth-context";
import { CreateButton, GridWrapper } from "../../../TimelineList.styled";

const ExpandedCard = ({ note, randomColorElement }) => {
  const auth = useAuth();

  const likeNote = () => {
    if (auth.isAuthenticated) console.log("authenticated");
    else console.log("not authenticated");
  };

  const shareNote = () => {
    if (auth.isAuthenticated) console.log("authenticated");
    else console.log("not authenticated");
  };

  return (
    <>
      <NotesCardContainer>
        <Title variant='h7' component='div' textAlign='center'>
          {note.Title}
        </Title>
        <CardContent>
          <AvatarBackground randomColor={randomColorElement}>
            <AvatarContainer>
              <CustomAvatar variant='square'>N</CustomAvatar>
            </AvatarContainer>
            <StyledLink to='/404'>
              <AvatarUsernames>
                <AvatarAuthor>{note.Author}</AvatarAuthor>
              </AvatarUsernames>
            </StyledLink>
          </AvatarBackground>
          <ContentContainer sx={{ marginTop: "10px" }}>
            <Typography color='textSecondary'>{note.Content}</Typography>
          </ContentContainer>
          <ContentContainer sx={{ marginTop: "10px" }}>
            <Tags tags={note.Tags}></Tags>
          </ContentContainer>
          <Typography color='textSecondary' textAlign='center'>
            {note.PublishedDate}
          </Typography>
        </CardContent>
      </NotesCardContainer>
      <GridWrapper
        sx={{
          display: "flex",
          justifyContent: "space-between",
          width: "100%",
        }}
      >
        <CreateButton
          onClick={likeNote}
          sx={{ marginLeft: "15px", marginTop: "20px" }}
        >
          Curtir
        </CreateButton>
        <CreateButton
          onClick={shareNote}
          sx={{ marginRight: "15px", marginTop: "20px" }}
        >
          Compartilhar
        </CreateButton>
      </GridWrapper>
      <TextField
        label='Comentar'
        variant='standard'
        sx={{
          display: "flex",
          margin: "auto",
          marginTop: "20px",
          width: "90%",
        }}
      />
    </>
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
  randomColorElement: PropTypes.func.isRequired,
};

export default ExpandedCard;
