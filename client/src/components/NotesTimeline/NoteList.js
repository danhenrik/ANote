import PropTypes from "prop-types";
import { Box, ButtonBase, Grid } from "@mui/material";
import NoteCard from "./NoteCard";

const containerStyle = {
  display: "flex",
  justifyContent: "center",
  width: "70%",
  margin: "0 auto",
};
function NoteList({ notes }) {
  return (
    <Box style={containerStyle}>
      <Grid container spacing={2}>
        {notes.map((note) => (
          <Grid item xs={12} sm={8} md={4} key={note.Id}>
            <ButtonBase sx={{ textAlign: "none" }}>
              <NoteCard note={note} />
            </ButtonBase>
          </Grid>
        ))}
      </Grid>
    </Box>
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
