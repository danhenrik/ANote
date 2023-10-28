import { useEffect, useState } from "react";
import CommunityList from "./CommunityList/CommunityList";
import { useAuth } from "../../../store/auth-context";
import useCommunities from "../../../api/useCommunities";
import { useSearchParams } from "react-router-dom";

const Communities = () => {
  const [communities, setCommunities] = useState([]);
  const userAuth = useAuth();
  const communitiesApi = useCommunities();
  const [searchParams] = useSearchParams();

  useEffect(() => {
    let fetchedCommunities = [];
    const fetchAndSetCommunities = async () => {
      if (userAuth.isAuthenticated) {
        if (!searchParams.get("world") || searchParams.get("world") === false) {
          fetchedCommunities = await communitiesApi.fetchCommunitiesByUser();
        } else {
          fetchedCommunities = await communitiesApi.fetchCommunities();
        }
      } else {
        fetchedCommunities = await communitiesApi.fetchCommunities();
      }

      setCommunities(fetchedCommunities);
    };

    fetchAndSetCommunities();
  }, [searchParams.get("world")]);

  return <CommunityList communities={communities}></CommunityList>;
};

export default Communities;
