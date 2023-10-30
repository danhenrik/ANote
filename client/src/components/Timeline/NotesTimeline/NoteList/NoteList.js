import PropTypes from "prop-types";
import { Button, ButtonBase, Grid } from "@mui/material";
import NoteCard from "../NoteCard/NoteCard";
import { useModal } from "../../../../store/modal-context";
import { useAuth } from "../../../../store/auth-context";
import NoteForm from "../NoteForm/NoteForm";
import LoginForm from "../../../AccessControl/Login/LoginForm";
import TimelineList from "../../TimelineList";
import useCommunities from "../../../../api/useCommunities";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import EmptyNotes from "../EmptyNotes";

const NoteList = ({
  notes,
  communityId,
  setNotesHandler,
  deleteNotesHandler,
}) => {
  const modal = useModal();
  const auth = useAuth();
  const communitiesApi = useCommunities();
  const [isFollowing, setIsFollowing] = useState();
  const { id } = useParams();

  useEffect(() => {
    if (auth.isAuthenticated) {
      const userFollowedCommunities = async () => {
        setIsFollowing(false);
        const communities = await communitiesApi.fetchCommunitiesByUser();
        if (communities) {
          if (communities.some((community) => community.Id === communityId)) {
            setIsFollowing(true);
          }
        }
      };
      userFollowedCommunities();
    }
  }, [useParams().id]);

  const followCommunity = async () => {
    const response = communitiesApi.followCommunity(communityId);
    if (response) {
      setIsFollowing(true);
    }
  };

  const handleAddNoteModal = () => {
    if (auth.isAuthenticated) {
      if (isFollowing || !id) {
        modal.openModal(
          <NoteForm
            notes={notes}
            communityId={communityId}
            closeModal={modal.closeModal}
            setNotesHandler={setNotesHandler}
          ></NoteForm>
        );
      } else {
        followCommunity(communityId);
      }
    } else {
      modal.openModal(<LoginForm closeModal={modal.closeModal}></LoginForm>);
    }
  };
  //trocar botao para um de seguir e um de adicionar, deve aparecer adicionar para usuario deslogado
  const buttonText =
    useParams().id && auth.isAuthenticated && isFollowing === false
      ? "Seguir Comunidade"
      : "Adicionar Nota";

  return (
    <TimelineList
      handleAddModal={handleAddNoteModal}
      addButtonText={buttonText}
    >
      {notes && notes.length ? (
        notes.map((note) => (
          <Grid item key={note.Id}>
            <ButtonBase>
              <NoteCard deleteNoteHandler={deleteNotesHandler} note={note} />
            </ButtonBase>
          </Grid>
        ))
      ) : (
        <EmptyNotes clickHandler={handleAddNoteModal}>
          Nenhuma Nota Aqui, Mas Você Pode Adicionar Uma!
        </EmptyNotes>
      )}
    </TimelineList>
  );
};

const noteShape = PropTypes.shape({
  Id: PropTypes.string.isRequired,
  Title: PropTypes.string.isRequired,
  Content: PropTypes.string.isRequired,
  LikeCount: PropTypes.number.isRequired,
  //Likes: PropTypes.arrayOf(PropTypes.string).isRequired,
  PublishedDate: PropTypes.string.isRequired,
  UpdatedDate: PropTypes.string,
  Author: PropTypes.string.isRequired,
  Tags: PropTypes.arrayOf(PropTypes.string),
  Communities: PropTypes.arrayOf(PropTypes.string),
  CommentCount: PropTypes.number.isRequired,
  //Commenters: PropTypes.arrayOf(PropTypes.string).isRequired,
});

NoteList.propTypes = {
  notes: PropTypes.arrayOf(noteShape),
  communityId: PropTypes.any,
  setNotesHandler: PropTypes.func.isRequired,
  deleteNotesHandler: PropTypes.func.isRequired,
};

export default NoteList;
