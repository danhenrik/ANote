import useApi from "./useApi";

const PAGE_SIZE = 8;

const handlePagination = (data, page, pageSize) => {
  const startIndex = (page - 1) * pageSize;
  const endIndex = startIndex + pageSize;
  return data.slice(startIndex, endIndex);
};

const mapApiNotesData = (data) => {
  return data.map((item) => ({
    Id: item.id,
    Title: item.title,
    Content: item.content,
    PublishedDate: item.created_at,
    UpdatedDate: item.updated_at,
    Author: item.author_id,
    Tags: item.tags.map((tag) => tag.name),
    Communities: item.communities.map((community) => community.name),
  }));
};

const fetchNotesFeedRequest = async (api, page) => {
  try {
    const response = await api.get("/notes/feed", {
      params: {
        page: 1,
        size: 1,
        sort_by: "title",
      },
    });
    return mapApiNotesData(handlePagination(response.data, page, PAGE_SIZE));
  } catch (error) {
    console.error("Error fetching notes:", error);
    throw error;
  }
};

const fetchNotesRequest = async (api, page) => {
  try {
    const response = await api.get("/notes", {
      params: {
        page: 1,
        size: 1,
        sort_by: "title",
      },
    });
    return mapApiNotesData(handlePagination(response.data, page, PAGE_SIZE));
  } catch (error) {
    console.error("Error fetching notes:", error);
    throw error;
  }
};

const fetchNotesByAuthorRequest = async (api, id) => {
  try {
    const response = await api.get(`/notes/author/${id}`);
    return mapApiNotesData(response.data.data);
  } catch (error) {
    console.error("Error fetching notes:", error);
    throw error;
  }
};

const fetchNotesByCommunityRequest = async (api, id) => {
  try {
    const response = await api.get(`/notes/community/${id}`);
    return mapApiNotesData(response.data.data);
  } catch (error) {
    console.error("Error fetching notes:", error);
    throw error;
  }
};

const createNoteRequest = async (api, note) => {
  try {
    const response = await api.post("notes", note);
    return response;
  } catch (error) {
    console.error("Error creating note:", error);
    throw error;
  }
};

const useNotes = () => {
  const api = useApi();

  const fetchNotesFeed = (page) => {
    return fetchNotesFeedRequest(api, page);
  };
  const fetchNotes = (page) => {
    return fetchNotesRequest(api, page);
  };

  const fetchNotesByAuthor = (id) => {
    return fetchNotesByAuthorRequest(api, id);
  };

  const fetchNotesByCommunity = (id) => {
    return fetchNotesByCommunityRequest(api, id);
  };

  const createNote = (note) => {
    return createNoteRequest(api, note);
  };

  return {
    fetchNotesFeed,
    fetchNotes,
    createNote,
    fetchNotesByAuthor,
    fetchNotesByCommunity,
  };
};

export default useNotes;
