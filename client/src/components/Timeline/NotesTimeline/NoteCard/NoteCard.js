import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import PropTypes from "prop-types";
import {
  AvatarAuthor,
  AvatarBackground,
  AvatarContainer,
  AvatarUsernames,
  ContentContainer,
  CustomAvatar,
  ModalStyling,
  NotesCardContainer,
  StyledLink,
  Title,
} from "./NoteCard.styled";
import Tags from "../../Tags/TagsList";
import { useState } from "react";
import { useModal } from "../../../../store/modal-context";
import ExpandedCard from "./ExpandedCard/ExpandedCard";
import LikeButton from "./LikeButton";
import CommentIcon from "@mui/icons-material/Comment";

const NoteCard = ({ note }) => {
  var randomColor = require("randomcolor");
  const modal = useModal();

  const [randomColorElement] = useState(
    randomColor({ luminosity: "light", format: "rgb" })
  );

  const handleExpandedNote = (event) => {
    if (event.target.closest("a")) return;
    if (event.target.closest("#like-button")) return;
    modal.openModal(
      <ExpandedCard
        note={note}
        randomColorElement={randomColorElement}
      ></ExpandedCard>
    );
    modal.setModalStyling(ModalStyling);
  };

  return (
    <NotesCardContainer
      onClick={handleExpandedNote}
      sx={{ minWidth: "300px", maxWidth: "300px" }}
    >
      <CardContent>
        <Title variant='h7' component='div'>
          {note.Title}
        </Title>
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
        <ContentContainer>
          <Typography color='textSecondary'>{note.Content}</Typography>
        </ContentContainer>
        <ContentContainer sx={{ marginTop: "10px" }}>
          <Tags tags={note.Tags}></Tags>
        </ContentContainer>
        <Typography color='textSecondary'>{note.PublishedDate}</Typography>
        <div
          id='note-actions'
          style={{ float: "right", marginBottom: "10px", marginTop: "10px" }}
        >
          <span id='like-button'>
            <LikeButton note={note}></LikeButton>
          </span>
          <CommentIcon style={{ marginLeft: "10px", color: "blue" }} />
        </div>
      </CardContent>
    </NotesCardContainer>
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

NoteCard.propTypes = {
  note: noteShape.isRequired,
};

export default NoteCard;
