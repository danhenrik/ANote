import PropTypes from "prop-types";
import { ButtonBase, Grid } from "@mui/material";
import NoteCard from "../NoteCard/NoteCard";
import {
  CreateButton as NotesButton,
  GridWrapper as NotesWrapper,
} from "./NoteList.styled";
import { useModal } from "../../../../store/modal-context";
import { useAuth } from "../../../../store/auth-context";
import NoteForm from "../NoteForm/NoteForm";
import LoginForm from "../../../AccessControl/Login/LoginForm";

function NoteList({ notes }) {
  const modal = useModal();
  const auth = useAuth();

  const handleAddNoteModal = () => {
    modal.openModal(
      auth.isAuthenticated ? (
        <NoteForm notes={notes} closeModal={modal.closeModal}></NoteForm>
      ) : (
        <LoginForm closeModal={modal.closeModal}></LoginForm>
      )
    );
  };

  return (
    <>
      <NotesWrapper>
        <Grid
          container
          spacing={{ xs: 2, md: 3 }}
          columns={{ xs: 1, sm: 1, md: 3 }}
        >
          <Grid
            item
            key='button'
            sx={{ display: "flex", justifyContent: "flex-end", width: "100%" }}
          >
            <NotesButton onClick={handleAddNoteModal}>
              Adicionar Nota
            </NotesButton>
          </Grid>
          {notes.map((note) => (
            <Grid item key={note.Id}>
              <ButtonBase>
                <NoteCard note={note} />
              </ButtonBase>
            </Grid>
          ))}
        </Grid>
      </NotesWrapper>
    </>
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

NoteList.propTypes = {
  notes: PropTypes.arrayOf(noteShape).isRequired,
};

export default NoteList;
