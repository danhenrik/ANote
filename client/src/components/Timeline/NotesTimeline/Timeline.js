import { useEffect, useState } from "react";
import NoteList from "./NoteList/NoteList";
import useNotes from "../../../api/useNotes";
import { useAuth } from "../../../store/auth-context";
import { useParams, useSearchParams } from "react-router-dom";

const Timeline = () => {
  const notesApi = useNotes();
  const [notes, setNotes] = useState([]);
  const userAuth = useAuth();
  const params = useParams();
  const [searchParams] = useSearchParams();

  const fetchAndSetNotes = async () => {
    let fetchedNotes = [];
    if (userAuth.isAuthenticated) {
      if (searchParams.get("world") && searchParams.get("world") == "true") {
        fetchedNotes = await notesApi.fetchNotesFeed();
      } else {
        if (params.id) {
          fetchedNotes = await notesApi.fetchNotesByCommunity(params.id);
        } else {
          fetchedNotes = await notesApi.fetchNotesByAuthor(
            userAuth.user.username
          );
        }
      }
    } else {
      setNotes([]);
      fetchedNotes = await notesApi.fetchNotes();
    }
    setNotes(fetchedNotes);
  };

  const setNotesHandler = (notes) => {
    fetchAndSetNotes(notes);
  };

  useEffect(() => {
    fetchAndSetNotes();
  }, [userAuth.isAuthenticated, searchParams.get("world"), params.id]);

  return (
    <NoteList
      communityId={params.id}
      setNotesHandler={setNotesHandler}
      notes={notes}
    ></NoteList>
  );
};

export default Timeline;
