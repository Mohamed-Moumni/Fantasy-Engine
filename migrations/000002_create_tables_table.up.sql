CREATE TABLE IF NOT EXISTS tables (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    team_id BIGINT NOT NULL,
    position INTEGER NOT NULL,
    played INTEGER NOT NULL DEFAULT 0,
    wins INTEGER NOT NULL DEFAULT 0,
    losses INTEGER NOT NULL DEFAULT 0,
    scores_for INTEGER NOT NULL DEFAULT 0,
    scores_against INTEGER NOT NULL DEFAULT 0,
    draws INTEGER NOT NULL DEFAULT 0,
    points INTEGER NOT NULL DEFAULT 0,
    score_diff_formatted INTEGER NOT NULL DEFAULT 0,
    
    CONSTRAINT fk_tables_team FOREIGN KEY (team_id) 
        REFERENCES teams(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_tables_team_id ON tables(team_id);
CREATE INDEX IF NOT EXISTS idx_tables_position ON tables(position);

COMMENT ON TABLE tables IS 'Team standings in leagues';
COMMENT ON COLUMN tables.score_diff_formatted IS 'Goal difference (scores_for - scores_against)';