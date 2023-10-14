import styled from "@emotion/styled";
import { Link } from "react-router-dom";

const TagList = styled("div")`
  display: flex;
  flex-wrap: wrap;
`;

const TagLink = styled(Link)`
  text-decoration: none;
  background-color: #7f56d9;
  color: #fff;
  border: none;
  border-radius: 4px;
  margin: 4px;
  padding: 8px 18px;
  cursor: pointer;

  &:hover {
    text-decoration: underline;
  }
`;

export { TagLink, TagList };
