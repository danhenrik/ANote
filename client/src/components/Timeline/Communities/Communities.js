import { useEffect, useState } from "react";
import CommunityList from "./CommunityList/CommunityList";
import myData from "./communities.json";
const Communities = () => {
  const [communities, setCommunities] = useState(myData);

  useEffect(() => {
    setCommunities(myData);
  }, [myData]);

  return <CommunityList communities={communities}></CommunityList>;
};

export default Communities;
