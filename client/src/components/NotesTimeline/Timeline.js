import { useState } from "react";
import myData from "./notes.json";
import NoteList from "./NoteList";

function Timeline() {
  const [notes] = useState(myData);
  return <NoteList notes={notes}></NoteList>;
}

export default Timeline;
