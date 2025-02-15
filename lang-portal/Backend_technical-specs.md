# Backend Server Technical Specs

## Project Overview

A language education institution seeks to develop a prototype learning management system that serves three primary functions:
- A comprehensive vocabulary database for language learning materials
- A performance tracking system that monitors and records student practice results
- A centralized platform for accessing various learning applications

## Technical Stack Requirements

- Backend Implementation: Go programming language
- Data Storage: SQLite database
- Web Framework: Gin for RESTful API development
- Response Format: All endpoints return JSON data
- Security Model: Open access (no authentication required)

## Database Structure

Our Database will be a single sqlite database called `words.db` that will be in the root of the project folder of `backend-go`.

The system utilizes the following data tables:
- words
    - id integer (primary key)
    - spanish string (target language text)
    - english string (translation)
    - parts json (grammatical information)

- words_groups
    - id integer (primary key)
    - word_id integer (foreign key to words)
    - group_id integer (foreign key to groups)
    
- groups
    - id integer (primary key)
    - name string (group identifier)

- study_sessions
    - id integer (primary key)
    - group_id integer (foreign key to groups)
    
- study_activities
    - id integer (primary key)
    - study_session_id integer (foreign key to study_sessions)
    - group_id integer (foreign key to groups)

- word_review_items
    - word_id integer (foreign key to words)
    - study_session_id integer (foreign key to study_sessions)
    - correct boolean (practice result)
    - created_at datetime (timestamp)

## API Endpoints

### Dashboard Endpoints
Collection of endpoints that provide overview statistics and recent activity data for the dashboard interface.

#### GET `/api/dashboard/last_study_session`
Retrieves details of the most recently completed study session, including timing and group information.
```json
{
  "id": 456,
  "group_id": 789,
  "created_at": "2024-03-15T09:45:12-05:00",
  "study_activity_id": 234,
  "group_name": "Common Verbs"
}
```

#### GET `/api/dashboard/study_progress`
Calculates and returns the overall learning progress metrics across all available vocabulary.
```json
{
  "total_words_studied": 25,
  "total_available_words": 250
}
```

#### GET `/api/dashboard/quick-stats`
Provides aggregated performance metrics and engagement statistics for the current user.
```json
{
  "success_rate": 75.5,
  "total_study_sessions": 12,
  "total_active_groups": 5,
  "study_streak_days": 7
}
```

### Study Activities Endpoints
Endpoints for managing and accessing various learning activities available in the system.

#### GET `/api/study_activities`
Lists all available study activities with their basic information and access URLs.
```json
[
  {
    "id": 3,
    "name": "Vocabulary Matching",
    "thumbnail_url": "/vocab-match.jpg",
    "launch_url": "/study-activities/3/launch"
  }
]
```

#### GET `/api/study_activities/:id`
Fetches detailed information about a specific study activity, including its description and visual assets.
```json
{
  "id": 3,
  "name": "Vocabulary Matching",
  "thumbnail_url": "https://example.com/vocab-match.jpg",
  "description": "Match Spanish words with their meanings"
}
```

#### GET `/api/study_activities/:id/study_sessions`
Retrieves a list of study sessions associated with a specific activity.
```json
{
  "items": [
    {
      "id": 789,
      "activity_name": "Vocabulary Matching",
      "group_name": "Common Verbs",
      "start_time": "2024-03-15T09:45:12-05:00",
      "end_time": "2024-03-15T10:15:12-05:00",
      "review_items_count": 15
    }
  ],
  "pagination": {
    "current_page": 2,
    "total_pages": 8,
    "total_items": 150,
    "items_per_page": 20
  }
}
```

#### POST `/api/study_activities`
Creates a new study activity entry.
```json
{
  "id": 567,
  "group_id": 789
}
```

### Words Endpoints
Endpoints for accessing and managing vocabulary items in the system.

#### GET `/api/words`
Retrieves a paginated list of vocabulary words with their translations and study statistics.
```json
{
  "items": [
    {
      "spanish": "comer",
      "english": "to eat",
      "correct_count": 12,
      "wrong_count": 3
    }
  ],
  "pagination": {
    "current_page": 3,
    "total_pages": 10,
    "total_items": 1000,
    "items_per_page": 100
  }
}
```

#### GET `/api/words/:id`
Provides comprehensive information about a specific word, including its associated groups and study statistics.
```json
{
  "spanish": "comer",
  "english": "to eat",
  "stats": {
    "correct_count": 12,
    "wrong_count": 3
  },
  "groups": [
    {
      "id": 4,
      "name": "Common Verbs"
    }
  ]
}
```

### Groups Endpoints
Endpoints for managing and accessing word groups and their associated study sessions.

#### GET `/api/groups`
Returns a paginated list of all vocabulary groups with their basic statistics.
```json
{
  "items": [
    {
      "id": 4,
      "name": "Common Verbs",
      "word_count": 50
    }
  ],
  "pagination": {
    "current_page": 2,
    "total_pages": 3,
    "total_items": 25,
    "items_per_page": 100
  }
}
```

#### GET `/api/groups/:id`
Fetches detailed information about a specific group, including its total word count.
```json
{
  "id": 4,
  "name": "Common Verbs",
  "stats": {
    "total_word_count": 50
  }
}
```

#### GET `/api/groups/:id/words`
Lists all words associated with a specific group.
```json
{
  "items": [
    {
      "spanish": "comer",
      "english": "to eat",
      "correct_count": 12,
      "wrong_count": 3
    }
  ],
  "pagination": {
    "current_page": 2,
    "total_pages": 3,
    "total_items": 50,
    "items_per_page": 100
  }
}
```

#### GET `/api/groups/:id/study_sessions`
Retrieves study sessions associated with a specific group.
```json
{
  "items": [
    {
      "id": 789,
      "activity_name": "Vocabulary Matching",
      "group_name": "Common Verbs",
      "start_time": "2024-03-15T09:45:12-05:00",
      "end_time": "2024-03-15T10:15:12-05:00",
      "review_items_count": 15
    }
  ],
  "pagination": {
    "current_page": 2,
    "total_pages": 3,
    "total_items": 45,
    "items_per_page": 100
  }
}
```

### Study Sessions Endpoints
Endpoints for tracking and managing individual study sessions and their results.

#### GET `/api/study_sessions`
Retrieves a chronological list of study sessions with their associated activities and metrics.
```json
{
  "items": [
    {
      "id": 789,
      "activity_name": "Vocabulary Matching",
      "group_name": "Common Verbs",
      "start_time": "2024-03-15T09:45:12-05:00",
      "end_time": "2024-03-15T10:15:12-05:00",
      "review_items_count": 15
    }
  ],
  "pagination": {
    "current_page": 3,
    "total_pages": 8,
    "total_items": 150,
    "items_per_page": 100
  }
}
```

#### GET `/api/study_sessions/:id`
Provides detailed information about a specific study session, including timing and performance data.
```json
{
  "id": 789,
  "activity_name": "Vocabulary Matching",
  "group_name": "Common Verbs",
  "start_time": "2024-03-15T09:45:12-05:00",
  "end_time": "2024-03-15T10:15:12-05:00",
  "review_items_count": 15
}
```

#### GET `/api/study_sessions/:id/words`
Lists all words reviewed during a specific study session with their results.
```json
{
  "items": [
    {
      "spanish": "comer",
      "english": "to eat",
      "correct_count": 12,
      "wrong_count": 3
    }
  ],
  "pagination": {
    "current_page": 2,
    "total_pages": 3,
    "total_items": 45,
    "items_per_page": 100
  }
}
```

### System Endpoints
Administrative endpoints for system maintenance and data management.

#### POST `/api/reset_history`
Clears all study session history while preserving word and group data.
```json
{
  "success": true,
  "message": "Successfully reset study history for all users"
}
```

#### POST `/api/full_reset`
Performs a complete system reset, reinitializing the database with seed data.
```json
{
  "success": true,
  "message": "Successfully reset system and reseeded database"
}
```

#### POST `/api/study_sessions/:id/words/:word_id/review`
Records the result of a word review during a study session, tracking correct/incorrect responses.
Request Payload:
```json
{
  "correct": false
}
```
Response:
```json
{
  "success": true,
  "word_id": 45,
  "study_session_id": 789,
  "correct": false,
  "created_at": "2024-03-15T10:12:33-05:00"
}
```

## Mage Tasks

Mage is a Make-like build tool using Go. It allows you to write your build scripts in Go, providing the benefits of Go's type safety and tooling.

Let's list out possible tasks we need for our lang portal.

### Initialize Database 
This task will initialize the sqlite database called `words.db`

### Migrate Database 
This task will run a series of migrations sql files on the database.

Migrations live in the `Migrations` folder.
The migration files will be run in order of their file name. 
The file names shoukd looks like this:

```sql
0001_init.sql
0002_create_words_table.sql
```

### Seed Database
This task will import json files and transform them into target data for our database.

All seed files live in the `seeds` folder

In our task we should have DSL to specify each seed file and its expected group word name.



```json
[
    {
        "spanish": "comer",
        "english": "to eat"
    }
]
