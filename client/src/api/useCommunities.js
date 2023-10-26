import useApi from "./useApi";

const mapApiCommunitiesData = (data) => {
  return data.map((item) => ({
    Id: item.id,
    Name: item.name,
    //Tags: item.tags.map((tag) => tag.name),
    Tags: ["abc"],
  }));
};

const fetchCommunitiesRequest = async (api) => {
  try {
    const response = await api.get("/communities");
    return mapApiCommunitiesData(response.data.data);
  } catch (error) {
    console.error("Error fetching communities:", error);
    throw error;
  }
};

const createCommunityRequest = async (api, community) => {
  try {
    const response = await api.post("communities", community);
    return mapApiCommunitiesData(response.data);
  } catch (error) {
    console.error("Error creating community:", error);
    throw error;
  }
};

const followCommunityRequest = async (api, community) => {
  try {
    const response = await api.post(`communities/join/, ${community}`);
    return response.data;
  } catch (error) {
    console.error("Error creating community:", error);
    throw error;
  }
};

const unfollowCommunityRequest = async (api, community) => {
  try {
    const response = await api.post(`communities/leave/, ${community}`);
    return response.data;
  } catch (error) {
    console.error("Error creating community:", error);
    throw error;
  }
};

const useCommunities = () => {
  const api = useApi();

  const fetchCommunities = () => {
    return fetchCommunitiesRequest(api);
  };

  const createCommunity = (community) => {
    return createCommunityRequest(api, community);
  };

  const followCommunity = (community) => {
    return followCommunityRequest(api, community);
  };

  const unfollowCommunity = (community) => {
    return unfollowCommunityRequest(api, community);
  };

  return {
    fetchCommunities,
    createCommunity,
    followCommunity,
    unfollowCommunity,
  };
};

export default useCommunities;
