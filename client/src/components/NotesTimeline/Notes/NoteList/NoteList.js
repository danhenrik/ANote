import PropTypes from "prop-types";
import { ButtonBase, Grid } from "@mui/material";
import NoteCard from "../NoteCard/NoteCard";
import { NotesContainer } from "./NoteList.styled";

function NoteList({ notes }) {
  return (
    <NotesContainer>
      <Grid container spacing={2}>
        {notes.map((note) => (
          <Grid item xs={12} sm={8} md={4} key={note.Id}>
            <ButtonBase sx={{ textAlign: "none" }}>
              <NoteCard note={note} />
            </ButtonBase>
          </Grid>
        ))}
      </Grid>
    </NotesContainer>
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
