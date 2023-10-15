import { useEffect, useState } from "react";
import myData from "./Notes/notes.json";
import NoteList from "./Notes/NoteList/NoteList";

function Timeline() {
  const [notes, setNotes] = useState(myData);

  useEffect(() => {
    setNotes(myData);
  }, [myData]);

  return (
    <div>
      <NoteList notes={notes}></NoteList>;
    </div>
  );
}

export default Timeline;
