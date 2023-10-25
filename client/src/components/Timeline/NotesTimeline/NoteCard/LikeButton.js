import React, { useState } from "react";
import FavoriteIcon from "@mui/icons-material/Favorite";
import FavoriteBorderIcon from "@mui/icons-material/FavoriteBorder";
import axios from "axios";

const LikeButton = (note) => {
  const [likes, setLikes] = useState(note.note.LikesCount);
  const [isClicked, setIsClicked] = useState(false);
  const [isHovered, setIsHovered] = useState(false);
  const handleClick = async () => {
    try {
      if (isClicked) {
        axios.delete("/likes/" + note.note.AuthorId + "/" + note.note.Id);
        setLikes(likes - 1);
      } else {
        const likeData = {
          user_id: note.note.AuthorId,
          note_id: note.note.Id,
        };
        axios.post("/likes", likeData);
        setLikes(likes + 1);
      }
      setIsClicked(!isClicked);
    } catch (error) {
      console.log("Like failed:", error);
    }
  };

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
            className={`like-button ${isClicked && "liked"}`}
            style={favoriteIconStyling}
          />
        ) : (
          <FavoriteBorderIcon
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

export default LikeButton;
