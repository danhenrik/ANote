import { useEffect, useState } from "react";
import NoteList from "./NoteList/NoteList";
import useNotes from "../../../api/useNotes";

const Timeline = () => {
  const notesApi = useNotes();
  const [notes, setNotes] = useState([]);

  useEffect(() => {
    const fetchAndSetNotes = async () => {
      const fetchedNotes = await notesApi.fetchNotes(1);
      setNotes(fetchedNotes);
    };

    fetchAndSetNotes();
  }, []);

  return <NoteList notes={notes}></NoteList>;
};

export default Timeline;
