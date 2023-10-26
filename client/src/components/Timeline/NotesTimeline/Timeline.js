import { useEffect, useState } from "react";
import NoteList from "./NoteList/NoteList";
import useNotes from "../../../api/useNotes";
import { useAuth } from "../../../store/auth-context";
import { useParams } from "react-router-dom";

const Timeline = () => {
  const notesApi = useNotes();
  const [notes, setNotes] = useState([]);
  const userAuth = useAuth();
  const params = useParams();

  useEffect(() => {
    let fetchedNotes = [];
    const fetchAndSetNotes = async () => {
      if (userAuth.isAuthenticated) {
        if (params.id) {
          fetchedNotes = await notesApi.fetchNotesByCommunity(params.id);
        } else {
          fetchedNotes = await notesApi.fetchNotesByAuthor(
            userAuth.user.username
          );
        }
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
