import useApi from "./useApi";

const mapApiCommunitiesData = async (data, api) => {
  const communitiesWithImages = await Promise.all(
    data.map(async (item) => {
      try {
        await api.get(`/static/${item.background}`);
        const imageUrl = `/static/${item.background}`;
        return {
          Id: item.id,
          Name: item.name,
          Background: imageUrl,
        };
      } catch (error) {
        console.error("Error fetching image:", error);
        return {
          Id: item.id,
          Name: item.name,
          Background: null,
        };
      }
    })
  );

  return communitiesWithImages;
};

const fetchCommunitiesRequest = async (api) => {
  try {
    const response = await api.get("/communities");
    return mapApiCommunitiesData(response.data.data, api);
  } catch (error) {
    console.error("Error fetching communities:", error);
  }
};

const fetchCommunitiesByUserRequest = async (api) => {
  try {
    const response = await api.get("/communities/my");
    return mapApiCommunitiesData(response.data.data, api);
  } catch (error) {
    console.error("Error fetching communities:", error);
  }
};

const createCommunityRequest = async (api, community) => {
  try {
    const response = await api.post("communities", community);
    return response.data;
  } catch (error) {
    console.error("Error creating community:", error);
  }
};

const followCommunityRequest = async (api, community) => {
  try {
    const response = await api.post(`communities/join/${community}`);
    return response.data;
  } catch (error) {
    console.error("Error creating community:", error);
  }
};

const unfollowCommunityRequest = async (api, community) => {
  try {
    const response = await api.post(`communities/leave/${community}`);
    return response.data;
  } catch (error) {
    console.error("Error creating community:", error);
  }
};

const useCommunities = () => {
  const api = useApi();

  const fetchCommunities = () => {
    return fetchCommunitiesRequest(api);
  };

  const fetchCommunitiesByUser = () => {
    return fetchCommunitiesByUserRequest(api);
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
    fetchCommunitiesByUser,
  };
};

export default useCommunities;
