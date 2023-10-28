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

  useEffect(() => {
    let fetchedNotes = [];
    const fetchAndSetNotes = async () => {
      if (userAuth.isAuthenticated) {
        if (searchParams.get("world") && searchParams.get("world") === true) {
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

    fetchAndSetNotes();
  }, [userAuth.isAuthenticated]);

  return <NoteList communityId={params.id} notes={notes}></NoteList>;
};

export default Timeline;
