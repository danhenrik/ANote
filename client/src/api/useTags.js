import useApi from "./useApi";

const useTags = () => {
  const api = useApi();

  const mapApiTagsData = (data) => {
    return data.map((item) => ({
      Id: item.id,
      Tags: item.name,
    }));
  };

  const fetchTags = async () => {
    try {
      const response = await api.get("/tags");
      return mapApiTagsData(response.data.data);
    } catch (error) {
      console.error("Error fetching tags:", error);
      throw error;
    }
  };

  return {
    fetchTags,
  };
};

export default useTags;
