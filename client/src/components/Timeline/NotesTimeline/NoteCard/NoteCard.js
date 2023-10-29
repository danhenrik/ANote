import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import PropTypes from "prop-types";
import {
  AvatarAuthor,
  AvatarBackground,
  AvatarContainer,
  AvatarUsernames,
  CommentButton,
  ContentContainer,
  CustomAvatar,
  ModalStyling,
  NotesCardContainer,
  StyledLink,
  Title,
} from "./NoteCard.styled";
import Tags from "../../Tags/TagsList";
import { useState, useEffect } from "react";
import { useModal } from "../../../../store/modal-context";
import ExpandedCard from "./ExpandedCard/ExpandedCard";
import LikeButton from "./LikeButton";
import formatDate from "../../../../util/formatDate";
import { useAuth } from "../../../../store/auth-context";
import axios from "axios";

const NoteCard = ({ note }) => {
  var randomColor = require("randomcolor");
  const modal = useModal();
  const userAuth = useAuth();

  const [randomColorElement] = useState(
    randomColor({ luminosity: "light", format: "rgb" })
  );
  const formatedDate = formatDate(note.PublishedDate);

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

  const [numberComments, setComments] = useState(0);
  useEffect(() => {
    const initComments = async () => {
      try {
        const comments = await axios.get("/comments/count/" + note.Id);
        setComments(comments.data.data);
      } catch (error) {
        console.log("Comment number retrieving failed: ", error);
      }
    };

    initComments();
  }, []);

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
          <AvatarUsernames>
            <AvatarAuthor>{note.Author}</AvatarAuthor>
          </AvatarUsernames>
        </AvatarBackground>
        <ContentContainer>
          <Typography color='textSecondary'>{note.Content}</Typography>
        </ContentContainer>
        <ContentContainer sx={{ marginTop: "10px" }}>
          <Tags tags={note.Tags}></Tags>
        </ContentContainer>

        {userAuth.isAuthenticated ? (
          <div
            id='note-actions'
            style={{
              display: "flex",
              justifyContent: "flex-end",
              marginTop: "10px",
            }}
          >
            <LikeButton sx={{ marginLeft: "8px" }} note={note}></LikeButton>
            <CommentButton />
            {
              <span
                style={{
                  color: "blue",
                  marginTop: "4px",
                  marginLeft: "5px",
                }}
              >
                {numberComments}
              </span>
            }
          </div>
        ) : (
          <></>
        )}
        <Typography color='textSecondary'>
          {formatedDate.day} às {formatedDate.hour}
        </Typography>
      </CardContent>
    </NotesCardContainer>
  );
};

const noteShape = PropTypes.shape({
  Id: PropTypes.string.isRequired,
  Title: PropTypes.string.isRequired,
  Content: PropTypes.string.isRequired,
  PublishedDate: PropTypes.string.isRequired,
  UpdatedDate: PropTypes.string.isRequired,
  Author: PropTypes.string.isRequired,
  Tags: PropTypes.arrayOf(PropTypes.string).isRequired,
});

NoteCard.propTypes = {
  note: noteShape.isRequired,
};

export default NoteCard;
