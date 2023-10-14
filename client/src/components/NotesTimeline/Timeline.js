import { useState } from "react";
import myData from "./Notes/notes.json";
import NoteList from "./Notes/NoteList/NoteList";
import { Button } from "@mui/material";
import { useModal } from "../../store/modal-context";
import NoteForm from "./Notes/NoteForm/NoteForm";
import { useAuth } from "../../store/auth-context";
import LoginForm from "../AccessControl/Login/LoginForm";

function Timeline() {
  const [notes] = useState(myData);
  const modal = useModal();
  const auth = useAuth();

  const handleAddNoteModal = () => {
    modal.openModal(
      auth.isAuthenticated ? (
        <NoteForm closeModal={modal.closeModal}></NoteForm>
      ) : (
        <LoginForm closeModal={modal.closeModal}></LoginForm>
      )
    );
  };

  return (
    <div>
      <Button onClick={handleAddNoteModal} sx={{ alignSelf: "right" }}>
        abc
      </Button>
      <NoteList notes={notes}></NoteList>;
    </div>
  );
}

export default Timeline;
