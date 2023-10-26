import { useEffect, useState } from "react";
import NoteList from "./NoteList/NoteList";
import useNotes from "../../../api/useNotes";
import { useAuth } from "../../../store/auth-context";

const Timeline = () => {
  const notesApi = useNotes();
  const [notes, setNotes] = useState([]);
  const user = useAuth();

  useEffect(() => {
    let fetchedNotes = [];
    const fetchAndSetNotes = async () => {
      if (user.isAuthenticated) {
        fetchedNotes = await notesApi.fetchNotesByAuthor(user.user.username);
      } else {
        fetchedNotes = await notesApi.fetchNotes(1);
      }

      setNotes(fetchedNotes);
    };

    fetchAndSetNotes();
  }, []);

  return <NoteList notes={notes}></NoteList>;
};

export default Timeline;
