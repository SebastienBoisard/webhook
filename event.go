package main

import "time"

type PushEvent struct {
	ObjectKind string `json:"object_kind"`
	Before string `json:"before"`
	After string `json:"after"`
	Ref string `json:"ref"`
	CheckoutSha string `json:"checkout_sha"`
	UserID int `json:"user_id"`
	UserName string `json:"user_name"`
	UserEmail string `json:"user_email"`
	UserAvatar string `json:"user_avatar"`
	ProjectID int `json:"project_id"`
	Project struct {
		Name string `json:"name"`
		Description string `json:"description"`
		WebURL string `json:"web_url"`
		AvatarURL interface{} `json:"avatar_url"`
		GitSSHURL string `json:"git_ssh_url"`
		GitHTTPURL string `json:"git_http_url"`
		Namespace string `json:"namespace"`
		VisibilityLevel int `json:"visibility_level"`
		PathWithNamespace string `json:"path_with_namespace"`
		DefaultBranch string `json:"default_branch"`
		Homepage string `json:"homepage"`
		URL string `json:"url"`
		SSHURL string `json:"ssh_url"`
		HTTPURL string `json:"http_url"`
	} `json:"project"`
	Repository struct {
		Name string `json:"name"`
		URL string `json:"url"`
		Description string `json:"description"`
		Homepage string `json:"homepage"`
		GitHTTPURL string `json:"git_http_url"`
		GitSSHURL string `json:"git_ssh_url"`
		VisibilityLevel int `json:"visibility_level"`
	} `json:"repository"`
	Commits []struct {
		ID string `json:"id"`
		Message string `json:"message"`
		Timestamp time.Time `json:"timestamp"`
		URL string `json:"url"`
		Author struct {
			Name string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
		Added []string `json:"added"`
		Modified []string `json:"modified"`
		Removed []interface{} `json:"removed"`
	} `json:"commits"`
	TotalCommitsCount int `json:"total_commits_count"`
}

type TagPushEvent struct {
	ObjectKind string `json:"object_kind"`
	Before string `json:"before"`
	After string `json:"after"`
	Ref string `json:"ref"`
	CheckoutSha string `json:"checkout_sha"`
	UserID int `json:"user_id"`
	UserName string `json:"user_name"`
	UserAvatar string `json:"user_avatar"`
	ProjectID int `json:"project_id"`
	Project struct {
		Name string `json:"name"`
		Description string `json:"description"`
		WebURL string `json:"web_url"`
		AvatarURL interface{} `json:"avatar_url"`
		GitSSHURL string `json:"git_ssh_url"`
		GitHTTPURL string `json:"git_http_url"`
		Namespace string `json:"namespace"`
		VisibilityLevel int `json:"visibility_level"`
		PathWithNamespace string `json:"path_with_namespace"`
		DefaultBranch string `json:"default_branch"`
		Homepage string `json:"homepage"`
		URL string `json:"url"`
		SSHURL string `json:"ssh_url"`
		HTTPURL string `json:"http_url"`
	} `json:"project"`
	Repository struct {
		Name string `json:"name"`
		URL string `json:"url"`
		Description string `json:"description"`
		Homepage string `json:"homepage"`
		GitHTTPURL string `json:"git_http_url"`
		GitSSHURL string `json:"git_ssh_url"`
		VisibilityLevel int `json:"visibility_level"`
	} `json:"repository"`
	Commits []interface{} `json:"commits"`
	TotalCommitsCount int `json:"total_commits_count"`
}


type IssueEvent struct {
	ObjectKind string `json:"object_kind"`
	User struct {
		Name string `json:"name"`
		Username string `json:"username"`
		AvatarURL string `json:"avatar_url"`
	} `json:"user"`
	Project struct {
		Name string `json:"name"`
		Description string `json:"description"`
		WebURL string `json:"web_url"`
		AvatarURL interface{} `json:"avatar_url"`
		GitSSHURL string `json:"git_ssh_url"`
		GitHTTPURL string `json:"git_http_url"`
		Namespace string `json:"namespace"`
		VisibilityLevel int `json:"visibility_level"`
		PathWithNamespace string `json:"path_with_namespace"`
		DefaultBranch string `json:"default_branch"`
		Homepage string `json:"homepage"`
		URL string `json:"url"`
		SSHURL string `json:"ssh_url"`
		HTTPURL string `json:"http_url"`
	} `json:"project"`
	Repository struct {
		Name string `json:"name"`
		URL string `json:"url"`
		Description string `json:"description"`
		Homepage string `json:"homepage"`
	} `json:"repository"`
	ObjectAttributes struct {
		ID int `json:"id"`
		Title string `json:"title"`
		AssigneeID int `json:"assignee_id"`
		AuthorID int `json:"author_id"`
		ProjectID int `json:"project_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Position int `json:"position"`
		BranchName interface{} `json:"branch_name"`
		Description string `json:"description"`
		MilestoneID interface{} `json:"milestone_id"`
		State string `json:"state"`
		Iid int `json:"iid"`
		URL string `json:"url"`
		Action string `json:"action"`
	} `json:"object_attributes"`
	Assignee struct {
		Name string `json:"name"`
		Username string `json:"username"`
		AvatarURL string `json:"avatar_url"`
	} `json:"assignee"`
}

type CommentEvent struct {
	ObjectKind string `json:"object_kind"`
	User struct {
		Name string `json:"name"`
		Username string `json:"username"`
		AvatarURL string `json:"avatar_url"`
	} `json:"user"`
	ProjectID int `json:"project_id"`
	Project struct {
		Name string `json:"name"`
		Description string `json:"description"`
		WebURL string `json:"web_url"`
		AvatarURL interface{} `json:"avatar_url"`
		GitSSHURL string `json:"git_ssh_url"`
		GitHTTPURL string `json:"git_http_url"`
		Namespace string `json:"namespace"`
		VisibilityLevel int `json:"visibility_level"`
		PathWithNamespace string `json:"path_with_namespace"`
		DefaultBranch string `json:"default_branch"`
		Homepage string `json:"homepage"`
		URL string `json:"url"`
		SSHURL string `json:"ssh_url"`
		HTTPURL string `json:"http_url"`
	} `json:"project"`
	Repository struct {
		Name string `json:"name"`
		URL string `json:"url"`
		Description string `json:"description"`
		Homepage string `json:"homepage"`
	} `json:"repository"`
	ObjectAttributes struct {
		ID int `json:"id"`
		Note string `json:"note"`
		NoteableType string `json:"noteable_type"`
		AuthorID int `json:"author_id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		ProjectID int `json:"project_id"`
		Attachment interface{} `json:"attachment"`
		LineCode string `json:"line_code"`
		CommitID string `json:"commit_id"`
		NoteableID interface{} `json:"noteable_id"`
		System bool `json:"system"`
		StDiff struct {
			Diff string `json:"diff"`
			NewPath string `json:"new_path"`
			OldPath string `json:"old_path"`
			AMode string `json:"a_mode"`
			BMode string `json:"b_mode"`
			NewFile bool `json:"new_file"`
			RenamedFile bool `json:"renamed_file"`
			DeletedFile bool `json:"deleted_file"`
		} `json:"st_diff"`
		URL string `json:"url"`
	} `json:"object_attributes"`
	Commit struct {
		ID string `json:"id"`
		Message string `json:"message"`
		Timestamp time.Time `json:"timestamp"`
		URL string `json:"url"`
		Author struct {
			Name string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
	} `json:"commit"`
}

type MergeRequestEvent struct {
	ObjectKind string `json:"object_kind"`
	User struct {
		Name string `json:"name"`
		Username string `json:"username"`
		AvatarURL string `json:"avatar_url"`
	} `json:"user"`
	ObjectAttributes struct {
		ID int `json:"id"`
		TargetBranch string `json:"target_branch"`
		SourceBranch string `json:"source_branch"`
		SourceProjectID int `json:"source_project_id"`
		AuthorID int `json:"author_id"`
		AssigneeID int `json:"assignee_id"`
		Title string `json:"title"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		StCommits interface{} `json:"st_commits"`
		StDiffs interface{} `json:"st_diffs"`
		MilestoneID interface{} `json:"milestone_id"`
		State string `json:"state"`
		MergeStatus string `json:"merge_status"`
		TargetProjectID int `json:"target_project_id"`
		Iid int `json:"iid"`
		Description string `json:"description"`
		Source struct {
			Name string `json:"name"`
			Description string `json:"description"`
			WebURL string `json:"web_url"`
			AvatarURL interface{} `json:"avatar_url"`
			GitSSHURL string `json:"git_ssh_url"`
			GitHTTPURL string `json:"git_http_url"`
			Namespace string `json:"namespace"`
			VisibilityLevel int `json:"visibility_level"`
			PathWithNamespace string `json:"path_with_namespace"`
			DefaultBranch string `json:"default_branch"`
			Homepage string `json:"homepage"`
			URL string `json:"url"`
			SSHURL string `json:"ssh_url"`
			HTTPURL string `json:"http_url"`
		} `json:"source"`
		Target struct {
			Name string `json:"name"`
			Description string `json:"description"`
			WebURL string `json:"web_url"`
			AvatarURL interface{} `json:"avatar_url"`
			GitSSHURL string `json:"git_ssh_url"`
			GitHTTPURL string `json:"git_http_url"`
			Namespace string `json:"namespace"`
			VisibilityLevel int `json:"visibility_level"`
			PathWithNamespace string `json:"path_with_namespace"`
			DefaultBranch string `json:"default_branch"`
			Homepage string `json:"homepage"`
			URL string `json:"url"`
			SSHURL string `json:"ssh_url"`
			HTTPURL string `json:"http_url"`
		} `json:"target"`
		LastCommit struct {
			ID string `json:"id"`
			Message string `json:"message"`
			Timestamp time.Time `json:"timestamp"`
			URL string `json:"url"`
			Author struct {
				Name string `json:"name"`
				Email string `json:"email"`
			} `json:"author"`
		} `json:"last_commit"`
		WorkInProgress bool `json:"work_in_progress"`
		URL string `json:"url"`
		Action string `json:"action"`
		Assignee struct {
			Name string `json:"name"`
			Username string `json:"username"`
			AvatarURL string `json:"avatar_url"`
		} `json:"assignee"`
	} `json:"object_attributes"`
}

type WikiPageEvent struct {
	ObjectKind string `json:"object_kind"`
	User struct {
		Name string `json:"name"`
		Username string `json:"username"`
		AvatarURL string `json:"avatar_url"`
	} `json:"user"`
	Project struct {
		Name string `json:"name"`
		Description string `json:"description"`
		WebURL string `json:"web_url"`
		AvatarURL interface{} `json:"avatar_url"`
		GitSSHURL string `json:"git_ssh_url"`
		GitHTTPURL string `json:"git_http_url"`
		Namespace string `json:"namespace"`
		VisibilityLevel int `json:"visibility_level"`
		PathWithNamespace string `json:"path_with_namespace"`
		DefaultBranch string `json:"default_branch"`
		Homepage string `json:"homepage"`
		URL string `json:"url"`
		SSHURL string `json:"ssh_url"`
		HTTPURL string `json:"http_url"`
	} `json:"project"`
	Wiki struct {
		WebURL string `json:"web_url"`
		GitSSHURL string `json:"git_ssh_url"`
		GitHTTPURL string `json:"git_http_url"`
		PathWithNamespace string `json:"path_with_namespace"`
		DefaultBranch string `json:"default_branch"`
	} `json:"wiki"`
	ObjectAttributes struct {
		Title string `json:"title"`
		Content string `json:"content"`
		Format string `json:"format"`
		Message string `json:"message"`
		Slug string `json:"slug"`
		URL string `json:"url"`
		Action string `json:"action"`
	} `json:"object_attributes"`
}

type PipelineEvent struct {
	ObjectKind string `json:"object_kind"`
	ObjectAttributes struct {
		ID int `json:"id"`
		Ref string `json:"ref"`
		Tag bool `json:"tag"`
		Sha string `json:"sha"`
		BeforeSha string `json:"before_sha"`
		Status string `json:"status"`
		Stages []string `json:"stages"`
		CreatedAt string `json:"created_at"`
		FinishedAt string `json:"finished_at"`
		Duration int `json:"duration"`
	} `json:"object_attributes"`
	User struct {
		Name string `json:"name"`
		Username string `json:"username"`
		AvatarURL string `json:"avatar_url"`
	} `json:"user"`
	Project struct {
		Name string `json:"name"`
		Description string `json:"description"`
		WebURL string `json:"web_url"`
		AvatarURL interface{} `json:"avatar_url"`
		GitSSHURL string `json:"git_ssh_url"`
		GitHTTPURL string `json:"git_http_url"`
		Namespace string `json:"namespace"`
		VisibilityLevel int `json:"visibility_level"`
		PathWithNamespace string `json:"path_with_namespace"`
		DefaultBranch string `json:"default_branch"`
	} `json:"project"`
	Commit struct {
		ID string `json:"id"`
		Message string `json:"message"`
		Timestamp time.Time `json:"timestamp"`
		URL string `json:"url"`
		Author struct {
			Name string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
	} `json:"commit"`
	Builds []struct {
		ID int `json:"id"`
		Stage string `json:"stage"`
		Name string `json:"name"`
		Status string `json:"status"`
		CreatedAt string `json:"created_at"`
		StartedAt interface{} `json:"started_at"`
		FinishedAt interface{} `json:"finished_at"`
		When string `json:"when"`
		Manual bool `json:"manual"`
		User struct {
			Name string `json:"name"`
			Username string `json:"username"`
			AvatarURL string `json:"avatar_url"`
		} `json:"user"`
		Runner interface{} `json:"runner"`
		ArtifactsFile struct {
			Filename interface{} `json:"filename"`
			Size interface{} `json:"size"`
		} `json:"artifacts_file"`
	} `json:"builds"`
}

type BuildEvent struct {
	ObjectKind string `json:"object_kind"`
	Ref string `json:"ref"`
	Tag bool `json:"tag"`
	BeforeSha string `json:"before_sha"`
	Sha string `json:"sha"`
	BuildID int `json:"build_id"`
	BuildName string `json:"build_name"`
	BuildStage string `json:"build_stage"`
	BuildStatus string `json:"build_status"`
	BuildStartedAt interface{} `json:"build_started_at"`
	BuildFinishedAt interface{} `json:"build_finished_at"`
	BuildDuration interface{} `json:"build_duration"`
	BuildAllowFailure bool `json:"build_allow_failure"`
	ProjectID int `json:"project_id"`
	ProjectName string `json:"project_name"`
	User struct {
		ID int `json:"id"`
		Name string `json:"name"`
		Email string `json:"email"`
	} `json:"user"`
	Commit struct {
		ID int `json:"id"`
		Sha string `json:"sha"`
		Message string `json:"message"`
		AuthorName string `json:"author_name"`
		AuthorEmail string `json:"author_email"`
		Status string `json:"status"`
		Duration interface{} `json:"duration"`
		StartedAt interface{} `json:"started_at"`
		FinishedAt interface{} `json:"finished_at"`
	} `json:"commit"`
	Repository struct {
		Name string `json:"name"`
		GitSSHURL string `json:"git_ssh_url"`
		Description string `json:"description"`
		Homepage string `json:"homepage"`
		GitHTTPURL string `json:"git_http_url"`
		VisibilityLevel int `json:"visibility_level"`
	} `json:"repository"`
}