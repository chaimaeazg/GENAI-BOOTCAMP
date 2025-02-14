# Backend Server Technical Specs

## Business Goal:

A language learning school wants to build a prototype of learning portal which will act as three things:
-Inventory of possible vocabulary that can be learned
- Act as a  Learning record store (LRS), providing correct and wrong score on practice vocabulary
- A unified launchpad to launch different learning apps


## Technical Requirements

- The Backend will be built using Go.
- The Database will be SQLite.
- The API will be built using Gin framework.
- The API will always return JSON
- There will be no Authentication 

## Database Schema

We have the following tables:
- words - stored vocabulary words
    - id integer
    - japanese string
    - romaji string
    - english string 
    - parts json

- words_groups - join table for words and groups many-to-many
    - id integer
    - word_id integer
    - group_id integer
    - 
- groups - thematic groups of words
    - id integer
    - name string
- study_sessions - record of study sessions grouping word_review_items
    - id integer
    - group_id integer
    
- study_activities - a specific study acivity, linking a study_session to group
    - id integer
    - study_session_id integer
    - group_id integer

- word_review_items - a record of word practice, determining if the word was correct or not 
    - word_id integer 
    - study_session_id integer
    - correct boolean
    - created_at datetime

## API Endpoints

- GET api/dashboard/last_study_session
- GET api/dashboard/study_progress
- GET api/dashboard/quick_stats
- GET api/study-activities/:id
- GET api/study-activities/:id/study-sessions
- POST api/study-activities/:id/launch
    - required params: group_id, study_activity_id
- GET api/words
    - pagination with 100 items per page
- GET api/words/:id
- GET api/groups 
    - pagination with 100 items per page
- GET api/groups/:id
- GET api/groups/:id/words
- GET api/groups/:id/study-sessions
- GET api/study-sessions
    - pagination with 100 items per page
- GET api/study-sessions/:id
- GET api/study-sessions/:id/words
- POST api/reset-history
- POST api/full-reset
- POST api/study-sessions/:id/words/:word_id/review
    - required params: correct