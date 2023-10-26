import { useEffect, useState } from "react";
import CommunityList from "./CommunityList/CommunityList";
import { useAuth } from "../../../store/auth-context";
import useCommunities from "../../../api/useCommunities";

const Communities = () => {
  const [communities, setCommunities] = useState([]);
  const userAuth = useAuth();
  const communitiesApi = useCommunities();

  useEffect(() => {
    let fetchedCommunities = [];
    const fetchAndSetCommunities = async () => {
      if (userAuth.isAuthenticated) {
        fetchedCommunities = await communitiesApi.fetchCommunities();
        /*
        fetchedCommunities = await communitiesApi.fetchCommunitiesByAuthor(
          userAuth.user.username
        );
        */
      } else {
        fetchedCommunities = await communitiesApi.fetchCommunities();
      }

      setCommunities(fetchedCommunities);
    };

    fetchAndSetCommunities();
  }, []);

  return <CommunityList communities={communities}></CommunityList>;
};

export default Communities;
