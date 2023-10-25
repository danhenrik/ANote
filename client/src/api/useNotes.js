import useApi from "./useApi";

const PAGE_SIZE = 8;

const handlePagination = (data, page, pageSize) => {
  const startIndex = (page - 1) * pageSize;
  const endIndex = startIndex + pageSize;
  return data.slice(startIndex, endIndex);
};

const fetchNotesRequest = async (api, page) => {
  try {
    const response = await api.get("notes");
    return handlePagination(response.data, page, PAGE_SIZE);
  } catch (error) {
    console.error("Error fetching notes:", error);
    throw error;
  }
};

const createNoteRequest = async (api, note) => {
  try {
    const response = await api.post("notes", note);
    return response.data;
  } catch (error) {
    console.error("Error creating note:", error);
    throw error;
  }
};

const useNotes = () => {
  const api = useApi();

  const fetchNotes = (page) => {
    return fetchNotesRequest(api, page);
  };

  const createNote = (note) => {
    return createNoteRequest(api, note);
  };

  return {
    fetchNotes,
    createNote,
  };
};

export default useNotes;
