import React, { useEffect, useState } from "react";
import NoteList from "./NoteList/NoteList";
import useNotes from "../../../api/useNotes";
import { useAuth } from "../../../store/auth-context";
import { useParams, useSearchParams } from "react-router-dom";
import { Box, Button } from "@mui/material";

const Timeline = () => {
  const notesApi = useNotes();
  const [notes, setNotes] = useState([]);
  const [page, setPage] = useState(1); // Initialize the page state with 1
  const userAuth = useAuth();
  const params = useParams();
  const [searchParams] = useSearchParams();

  function extractQueryParams(searchParams) {
    const queryParams = {};
    searchParams.forEach((value, key) => {
      if (key !== "search" && key !== "world") {
        queryParams[key] = value;
      }
    });

    return queryParams;
  }

  const fetchAndSetNotes = async () => {
    let fetchedNotes = [];
    if (searchParams.get("search") && searchParams.get("search") == "true") {
      const queryParams = extractQueryParams(searchParams);
      fetchedNotes = await notesApi.fetchNotesFilter(page, queryParams); // Use the 'page' state here
    } else {
      if (userAuth.isAuthenticated) {
        if (searchParams.get("world") && searchParams.get("world") == "true") {
          fetchedNotes = await notesApi.fetchNotes(page); // Use the 'page' state here
        } else {
          if (params.id) {
            fetchedNotes = await notesApi.fetchNotesByCommunity(
              page,
              params.id
            ); // Use the 'page' state here
          } else {
            fetchedNotes = await notesApi.fetchNotesFeed(page); // Use the 'page' state here
          }
        }
      } else {
        if (params.id) {
          fetchedNotes = await notesApi.fetchNotesByCommunity(page, params.id); // Use the 'page' state here
        } else {
          fetchedNotes = await notesApi.fetchNotes(page); // Use the 'page' state here
        }
      }

      setNotes(fetchedNotes);
    }
  };

  const setNotesHandler = (notes) => {
    fetchAndSetNotes(notes);
  };

  const deleteNotesHandler = (id) => {
    const updatedNotes = notes.filter((note) => note.Id !== id);
    setNotes(updatedNotes);
  };

  const handlePageChange = (newPage) => {
    setPage(newPage); // Update the 'page' state
  };

  useEffect(() => {
    fetchAndSetNotes();
  }, [userAuth.isAuthenticated, searchParams, params.id, page]);

  return (
    <div>
      <NoteList
        communityId={params.id}
        setNotesHandler={setNotesHandler}
        deleteNotesHandler={deleteNotesHandler}
        notes={notes}
      ></NoteList>
      <Box
        display='flex'
        justifyContent='center'
        alignItems='center'
        mt='auto' // This pushes the Box component to the bottom
      >
        <Button
          onClick={() => handlePageChange(page - 1)}
          disabled={page === 1}
        >
          Anterior
        </Button>
        <span className='page-number'>{page}</span>
        <Button
          disabled={notes && notes.length < 8}
          onClick={() => handlePageChange(page + 1)}
        >
          Seguinte
        </Button>
      </Box>
    </div>
  );
};

export default Timeline;
