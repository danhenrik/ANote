import { useState } from "react";
import myData from "./Notes/notes.json";
import NoteList from "./Notes/NoteList/NoteList";

function Timeline() {
  const [notes] = useState(myData);
  return <NoteList notes={notes}></NoteList>;
}

export default Timeline;
