import React, { useEffect, useState } from "react";
import FavoriteIcon from "@mui/icons-material/Favorite";
import FavoriteBorderIcon from "@mui/icons-material/FavoriteBorder";
import axios from "axios";
import { useAuth } from "../../../../store/auth-context";
import PropTypes from "prop-types";

const LikeButton = ({ note }) => {
  const [likes, setLikes] = useState(0);
  const [isClicked, setIsClicked] = useState(false);
  const [isHovered, setIsHovered] = useState(false);
  const auth = useAuth();

  const handleClick = async () => {
    try {
      if (isClicked) {
        const deletedLike = await axios.delete(
          "/likes/" + note.Author + "/" + note.Id
        );
        if (deletedLike) {
          setLikes(likes - 1);
        }
      } else {
        const likeData = {
          user_id: note.Author,
          note_id: note.Id,
        };
        const postedLike = axios.post("/likes", likeData);
        if (postedLike) {
          setLikes(likes + 1);
        }
      }
      setIsClicked(!isClicked);
    } catch (error) {
      console.log("Like failed:", error);
    }
  };

  const hasUserLiked = (note) => {
    const likeValues = Object.values(note.Likes);

    return likeValues.some((like) => like.user_id === auth.username);
  };

  useEffect(() => {
    const initLikes = async (note) => {
      const numberLikes = note.LikeCount;
      setLikes(numberLikes);
      if (hasUserLiked(note)) setIsClicked(true);
    };

    initLikes(note);
  }, []);

  const favoriteIconStyling = {
    color: "red",
    ...(isHovered && { transform: "scale(1.2)" }),
  };

  return (
    <>
      <span
        onClick={handleClick}
        onMouseEnter={() => setIsHovered(true)}
        onMouseLeave={() => setIsHovered(false)}
      >
        {isClicked || isHovered ? (
          <FavoriteIcon
            id='like-button'
            className={`like-button ${isClicked && "liked"}`}
            style={favoriteIconStyling}
          />
        ) : (
          <FavoriteBorderIcon
            id='like-button'
            className={`like-button ${isClicked && "liked"}`}
            style={{ marginRight: "2px", color: "red" }}
          />
        )}
      </span>
      <span
        style={{
          color: "red",
          marginTop: "4px",
        }}
      >
        {likes}
      </span>
    </>
  );
};

const noteShape = PropTypes.shape({
  Id: PropTypes.string.isRequired,
  Title: PropTypes.string.isRequired,
  Content: PropTypes.string.isRequired,
  PublishedDate: PropTypes.string.isRequired,
  UpdatedDate: PropTypes.string.isRequired,
  Author: PropTypes.string.isRequired,
  Tags: PropTypes.arrayOf(PropTypes.string).isRequired,
  CommentCount: PropTypes.number.isRequired,
  LikeCount: PropTypes.number.isRequired,
  Likes: PropTypes.arrayOf(PropTypes.any).isRequired,
});

LikeButton.propTypes = {
  note: noteShape.isRequired,
};

export default LikeButton;
