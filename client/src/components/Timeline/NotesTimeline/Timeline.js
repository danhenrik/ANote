import React, { useEffect, useState } from "react";
import NoteList from "./NoteList/NoteList";
import useNotes from "../../../api/useNotes";
import { useAuth } from "../../../store/auth-context";
import { useParams, useSearchParams } from "react-router-dom";
import { Box, Button } from "@mui/material";

const Timeline = () => {
  const notesApi = useNotes();
  const [notes, setNotes] = useState([]);
  const [page, setPage] = useState(1);
  const userAuth = useAuth();
  const params = useParams();
  const [searchParams] = useSearchParams();
  const [displayText, setDisplayText] = useState("Feed de Notas");

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
      fetchedNotes = await notesApi.fetchNotesFilter(
        page,
        queryParams,
        userAuth
      );
      setNotes(fetchedNotes);
      setDisplayText("Notas Populares");
    } else {
      if (userAuth.isAuthenticated) {
        if (searchParams.get("world") && searchParams.get("world") == "true") {
          setDisplayText("Notas Populares");
          fetchedNotes = await notesApi.fetchNotes(page, userAuth);
        } else {
          if (params.id) {
            setDisplayText("Notas da Comunidade: " + params.id);
            fetchedNotes = await notesApi.fetchNotesByCommunity(
              page,
              params.id
            );
          } else {
            setDisplayText("Feed de Notas");
            fetchedNotes = await notesApi.fetchNotesFeed(page, userAuth);
          }
        }
      } else {
        if (params.id) {
          setDisplayText("Notas da Comunidade: " + params.id);
          fetchedNotes = await notesApi.fetchNotesByCommunity(page, params.id);
        } else {
          setDisplayText("Notas Populares");
          fetchedNotes = await notesApi.fetchNotes(page, userAuth);
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
    setPage(newPage);
  };

  useEffect(() => {
    fetchAndSetNotes();
  }, [userAuth.isAuthenticated, searchParams, params.id, page]);

  return (
    <div>
      <NoteList
        communityId={searchParams.get("communityId")}
        setNotesHandler={setNotesHandler}
        deleteNotesHandler={deleteNotesHandler}
        notes={notes}
        displayText={displayText}
      ></NoteList>
      <Box
        style={{
          position: "fixed",
          bottom: 0,
          left: 0,
          width: "100%",
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          backgroundColor: "white",
          padding: "10px",
        }}
      >
        <Button
          onClick={() => handlePageChange(page - 1)}
          disabled={page === 1}
        >
          Anterior
        </Button>
        <span className='page-number'>{page}</span>
        <Button
          disabled={notes === undefined || notes.length < 8}
          onClick={() => handlePageChange(page + 1)}
        >
          Seguinte
        </Button>
      </Box>
    </div>
  );
};

export default Timeline;
