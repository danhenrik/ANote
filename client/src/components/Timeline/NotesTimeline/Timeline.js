import { useEffect, useState } from "react";
import myData from "./notes.json";
import NoteList from "./NoteList/NoteList";

const Timeline = () => {
  const [notes, setNotes] = useState(myData);

  useEffect(() => {
    setNotes(myData);
  }, [myData]);

  return <NoteList notes={notes}></NoteList>;
};

export default Timeline;
