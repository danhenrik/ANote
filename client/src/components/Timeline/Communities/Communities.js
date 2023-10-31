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
  const [displayText, setDisplayText] = useState("Comunidades Públicas");

  const fetchAndSetCommunities = async () => {
    let fetchedCommunities = [];
    if (userAuth.isAuthenticated) {
      if (!searchParams.get("world") || searchParams.get("world") === false) {
        setDisplayText("Minhas Comunidades");
        fetchedCommunities = await communitiesApi.fetchCommunitiesByUser();
      } else {
        setDisplayText("Comunidades Públicas");
        fetchedCommunities = await communitiesApi.fetchCommunities();
      }
    } else {
      setDisplayText("Comunidades Públicas");
      fetchedCommunities = await communitiesApi.fetchCommunities();
    }

    setCommunities(fetchedCommunities);
  };

  const setCommunitiesHandler = (communities) => {
    fetchAndSetCommunities(communities);
  };

  useEffect(() => {
    fetchAndSetCommunities();
  }, [searchParams.get("world"), userAuth.isAuthenticated]);

  return (
    <CommunityList
      displayText={displayText}
      communities={communities}
      setCommunitiesHandler={setCommunitiesHandler}
    ></CommunityList>
  );
};

export default Communities;
