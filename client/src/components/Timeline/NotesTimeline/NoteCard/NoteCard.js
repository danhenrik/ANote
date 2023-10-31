import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import PropTypes from "prop-types";
import DeleteIcon from "@mui/icons-material/Delete";
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
import { useParams } from "react-router-dom";
import { IconButton } from "@mui/material";
import useNotes from "../../../../api/useNotes";
import useApi from "../../../../api/useApi";

const NoteCard = ({ note, deleteNoteHandler }) => {
  var randomColor = require("randomcolor");
  const modal = useModal();
  const userAuth = useAuth();
  const isFollowing = useState(false);
  const notesApi = useNotes();
  const api = useApi();
  const [avatar, setAvatar] = useState("");

  const [randomColorElement] = useState(
    randomColor({ luminosity: "light", format: "rgb" })
  );
  const formatedDate = formatDate(note.PublishedDate);

  const handleExpandedNote = (event) => {
    if (event.target.closest("a")) return;
    if (event.target.closest("#like-button")) return;
    if (event.target.closest("#delete-button")) return;
    modal.openModal(
      <ExpandedCard
        note={note}
        numberComments={numberComments}
        numberCommentsHandler={setNumberComments}
        randomColorElement={randomColorElement}
        avatar={avatar}
      ></ExpandedCard>
    );
    modal.setModalStyling(ModalStyling);
  };

  const getAvatar = async (note) => {
    const response = await api.get("/users/username/" + note.Author);
    await api.get(`/static/${response.data.data.avatar}`);
    if (response) {
      setAvatar("/static/" + response.data.data.avatar);
    }
  };

  const [numberComments, setNumberComments] = useState(0);
  useEffect(() => {
    const initComments = async () => {
      const comments = note.CommentCount;
      setNumberComments(comments);
    };
    getAvatar(note);
    initComments();
  }, []);

  const deleteNote = async () => {
    const id = note.Id;
    const response = await notesApi.deleteNote(note.Id);
    if (response) {
      deleteNoteHandler(id);
    }
  };

  return (
    <NotesCardContainer
      onClick={handleExpandedNote}
      sx={{ minWidth: "300px", maxWidth: "300px" }}
    >
      <CardContent>
        {note.Author === userAuth.username && (
          <IconButton
            onClick={deleteNote}
            style={{
              position: "absolute",
              right: 0,
              top: 0,
              cursor: "pointer",
              color: "red",
              display: "block",
              transition: "color 0.3s, transform 0.3s",
            }}
          >
            <DeleteIcon id='delete-button' style={{ marginRight: "2px" }} />
          </IconButton>
        )}
        <Title
          variant='h7'
          component='div'
          style={{ textAlign: "center", textTransform: "uppercase" }}
        >
          {note.Title}
        </Title>
        <AvatarBackground
          randomColor={randomColorElement}
          style={{ marginTop: "10px" }}
        >
          <AvatarContainer>
            <CustomAvatar variant='square' src={avatar}></CustomAvatar>
          </AvatarContainer>
          <AvatarUsernames>
            <AvatarAuthor>{note.Author}</AvatarAuthor>
          </AvatarUsernames>
        </AvatarBackground>
        <ContentContainer style={{ marginTop: "10px" }}>
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
        {note.Communities && note.Communities.length > 0 && (
          <>
            <Typography color='black'>
              Comunidade: {note.Communities}
            </Typography>
          </>
        )}
        {note.Communities && note.Communities.length == 0 && (
          <>
            <Typography color='textSecondary'>Sem Comunidade</Typography>
          </>
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
  UpdatedDate: PropTypes.string,
  Author: PropTypes.string.isRequired,
  Tags: PropTypes.arrayOf(PropTypes.string).isRequired,
  CommentCount: PropTypes.number.isRequired,
  LikeCount: PropTypes.number.isRequired,
  Communities: PropTypes.any,
});

NoteCard.propTypes = {
  note: noteShape.isRequired,
  deleteNoteHandler: PropTypes.func.isRequired,
};

export default NoteCard;
