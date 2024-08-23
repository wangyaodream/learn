package main

import "multiple-services/service"


type userService struct {
    service.UnimplementedUsersServer
}

type repoService struct {
    service.UnimplementedRepoServer
}


