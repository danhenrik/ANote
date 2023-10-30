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
    PublishedDate: item.created_at ? item.created_at : item.published_date,
    UpdatedDate: item.updated_at ? item.update_at : item.updated_date,
    Author: item.author_id ? item.author_id : item.author,
    Tags: item.tags ? item.tags.map((tag) => tag.name) : [],
    Communities: item.communities
      ? item.communities.map((community) => community.name)
      : [],
    LikeCount: item.likes_count,
    CommentCount: item.comment_count,
    Likes: item.likes ? item.likes : [],
  }));
};

const fetchNotesFeedRequest = async (api, page) => {
  try {
    const response = await api.get("/notes/feed", {
      params: {
        page: page,
        size: PAGE_SIZE,
      },
    });
    if (response.data.data) {
      return handlePagination(
        mapApiNotesData(response.data.data),
        page,
        PAGE_SIZE
      );
    } else {
      return [];
    }
  } catch (error) {
    console.error("Error fetching notes:", error);
    throw error;
  }
};

const fetchNotesRequest = async (api, page) => {
  try {
    const response = await api.get("/notes", {
      params: {
        page: page,
        size: PAGE_SIZE,
      },
    });
    if (response.data.data) {
      return handlePagination(
        mapApiNotesData(response.data.data),
        page,
        PAGE_SIZE
      );
    } else {
      return [];
    }
  } catch (error) {
    console.error("Error fetching notes:", error);
    throw error;
  }
};

const fetchNotesByCommunityRequest = async (api, page, id) => {
  try {
    const response = await api.get("/notes", {
      params: {
        page: page,
        size: PAGE_SIZE,
        communities: id,
      },
    });
    if (response.data.data) {
      return handlePagination(
        mapApiNotesData(response.data.data),
        page,
        PAGE_SIZE
      );
    }
  } catch (error) {
    console.error("Error fetching notes:", error);
    throw error;
  }
};

const fetchNotesFilterRequest = async (api, page, filters) => {
  try {
    const response = await api.get("/notes", {
      params: {
        page: page,
        size: PAGE_SIZE,
        ...filters,
      },
    });
    if (response.data.data) {
      return handlePagination(
        mapApiNotesData(response.data.data),
        page,
        PAGE_SIZE
      );
    }
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

const deleteNoteRequest = async (api, noteId) => {
  try {
    const response = await api.delete(`/notes/${noteId}`);
    return response;
  } catch (error) {
    console.error(`Error deleting note with ID ${noteId}:`, error);
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

  const fetchNotesByCommunity = (page, id) => {
    return fetchNotesByCommunityRequest(api, page, id);
  };

  const fetchNotesFilter = (page, filters) => {
    return fetchNotesFilterRequest(api, page, filters);
  };

  const createNote = (note) => {
    return createNoteRequest(api, note);
  };

  const deleteNote = (noteId) => {
    return deleteNoteRequest(api, noteId);
  };

  return {
    fetchNotesFeed,
    fetchNotes,
    createNote,
    fetchNotesFilter,
    fetchNotesByCommunity,
    deleteNote,
  };
};

export default useNotes;
