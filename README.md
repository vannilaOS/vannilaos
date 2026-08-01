# Vanilla OS - מנהל חבילות Apx

Apx הוא מנהל החבילות ברירת המחדל של Vanilla OS. הוא עטיפה סביב מנהלי חבילות מרובים להתקנת חבילות והפעלת פקודות בתוך מיכל מנוהל.

## סקירה כללית

**Apx** (/à·peks/) מספק ממשק אחיד לניהול חבילות על פני הפצות Linux שונות. הוא משתמש ב-[distrobox](https://github.com/89luca89/distrobox) כדי ליצור סביבות בטחונות שבהן ניתן להתקין חבילות ללא השפעה על מערכת המארח שלך.

## תכונות

- **תמיכה בהפצות מרובות**: התקנת חבילות מהפצות Linux שונות
- **מבוסס על קונטיינר**: חבילות מותקנות בקונטיינרים מבודדים
- **הפשטת מנהל החבילות**: ממשק אחיד עבור מנהלי חבילות שונים (apt, dnf, pacman, וכו')
- **שילוב שולחן עבודה**: ייצוא אוטומטי של רשומות שולחן עבודה לאפליקציות מותקנות
- **ניהול תת-מערכות**: יצירה וניהול מרובות תת-מערכות עם מחסניות שונות

## הערימה

- **שפה**: Go 1.25.0
- **Framework**: Vanilla OS SDK + CLI builder
- **ממשק משתמש**: Charmbracelet (bubbletea, lipgloss) לממשק טרמינל
- **קונטיינר**: שילוב distrobox

## מבנה הפרויקט

```
cmd/
  main.go              # נקודת כניסה של היישום
  main_check.go        # וריאנט אימות מחרוזות
  locales/             # תמיכה בשפות מרובות (33+ שפות)

internal/cli/
  structs.go           # הגדרות פקודות CLI
  runtime.go           # מטפלי ביצוע תת-מערכות
  subsystems.go        # ניהול מחזור חיי תת-מערכות
  stacks.go            # פעולות ערימה
  pkgmanagers.go       # פעולות מנהל חבילות

core/                  # לוגיקה עסקית ליבית (תת-מערכות, ערימות, ניהול pkg)
config/                # קבצי תצורה
distrobox/             # סקריפטי שילוב distrobox
```

## בנייה

### דרישות מוקדמות

- Go 1.25.0 או גבוה יותר
- make
- podman או docker
- git

### שלבי בנייה

```bash
# שיבוט וכניסה לתיקייה
git clone --recursive https://github.com/yourusername/Vanilla-.git
cd Vanilla-

# בנייה
make build

# התקנה לרמת המערכת
sudo make install
sudo make install-manpages

# התקנה למיקום מותאם אישית
make install PREFIX=$HOME/.local
make install-manpages PREFIX=$HOME/.local
```

## שימוש

### פקודות בסיסיות

```bash
# רישום תת-מערכות זמינות
apx subsystems list

# יצירת תת-מערכת חדשה
apx subsystems new --name=ubuntu-latest

# כניסה לתת-מערכת
apx ubuntu-latest enter

# התקנת חבילות בתת-מערכת
apx ubuntu-latest install package-name

# הפעלת פקודה בתת-מערכת
apx ubuntu-latest run command

# רישום ערימות זמינות
apx stacks list

# ניהול מנהלי חבילות
apx pkgmanagers list
```

## גישה מרחוק

Apx תומך בגישה מרחוק דרך SSH, מה שמאפשר ניהול חבילות וביצוע פקודות בשרתים וציוד מרוחק.

### הגדרת SSH

```bash
# חיבור לשרת מרוחק
ssh user@remote-host

# הרצת פקודות Apx מרחוק
ssh user@remote-host "apx subsystems list"

# התקנת חבילות דרך SSH
ssh user@remote-host "apx ubuntu-latest install package-name"
```

### שימוש ב-iPhone או כלים ניידים

עבור גישה מרחוק מ-iPhone או מכשירים ניידים אחרים, תוכל להשתמש בתוכנות SSH כמו:

- **Termius** - ממשק SSH מלא עם תמיכה בקונטיינרים
- **iSH** - קונסולה SSH עבור iOS
- **Prompt** - קליינט SSH מקצועי לiOS

```bash
# דוגמה: חיבור מ-iPhone דרך Termius
ssh user@your-server.com
apx ubuntu-latest enter
```

### גישה מרחוק מאובטחת

```bash
# הגדרת מפתח SSH ציבורי לגישה מאובטחת
ssh-copy-id -i ~/.ssh/id_rsa.pub user@remote-host

# כניסה ללא סיסמה
ssh user@remote-host "apx subsystems list"

# הפעלת סקריפט מרחוק
ssh user@remote-host < /path/to/local/script.sh
```

## תמיכה ב-Android

Apx תומך בניהול חבילות וביצוע פקודות דרך מכשירי Android בעזרת Termux וSSH.

### Termux Integration

Termux היא טרמינל מלא עבור Android שתומכת בלימוד Apx בצורה ישירה.

```bash
# התקנת Termux מ-F-Droid
# https://f-droid.org/en/packages/com.termux/

# התקנת Apx ב-Termux
pkg install go git
git clone https://github.com/yourusername/Vanilla-.git
cd Vanilla-
make build
sudo make install
```

### גישה ל-Apx דרך Android

```bash
# חיבור לשרת Apx משרת האחסון שלך
ssh user@your-server.com

# הרצת פקודות Apx מרחוק
ssh user@your-server.com "apx subsystems list"

# התקנת חבילות מ-Android
ssh user@your-server.com "apx ubuntu-latest install python3"
```

### יישומי Android SSH

שימוש ביישומי SSH פופולריים עבור Android:

- **Termux** - שורת פקודה ישירה וSSH
- **Termius** - קליינט SSH מקצועי
- **JuiceSSH** - ממשק SSH משופר
- **ConnectBot** - קליינט SSH קל משקל

```bash
# דוגמה: חיבור דרך Termius
1. פתח Termius
2. צור חיבור SSH חדש
3. הזן את פרטי השרת שלך
4. הפעל: apx subsystems list
```

### פיתוח אפליקציות Android עם Apx

```bash
# התקנת כלי פיתוח Android בתת-מערכת
apx ubuntu-latest install android-sdk android-ndk

# קומפילציה של אפליקציות דרך SSH מ-Android
ssh user@dev-server.com "apx android-dev compile --project /path/to/project"

# בדיקה של עבודה עם ADB מרחוק
ssh user@dev-server.com "apx android-dev adb shell"
```

### סנכרון קבצים דרך Android

```bash
# העברת קבצים מ-Android לשרת Apx
scp /sdcard/myproject/* user@server.com:/home/user/projects/

# משיכת תוצאות מ-Apx ל-Android
scp user@server.com:/build/output/* /sdcard/Downloads/
```

## AI ויישומי בינה מלאכותית

Apx משתלב יכולות בינה מלאכותית כדי לשפר את חווית ניהול החבילות וההתקנה.

### המלצות חבילות בחכמה

Apx משתמש בלימוד מכונה כדי להמליץ על חבילות רלוונטיות על סמך:
- רישום התקנה קודם שלך
- תבניות שימוש נפוצות
- תלויויות ופרויקטים דומים

```bash
# קבל המלצות AI עבור חבילה
apx ai-recommend package-name

# ניתוח חבילה עם AI
apx ai-analyze ubuntu-latest
```

### פתרון בעיות בעזרת AI

AI יכול לסייע בטיפול בשגיאות וביצירת סביבה מאופטימלת:

```bash
# קבל אבחון AI לבעיות התקנה
apx ai-diagnose

# קבל הצעות AI לפתרון בעיות
apx ai-troubleshoot error-message

# קבל ייעוץ AI להגדרה אופטימלית
apx ai-optimize
```

### הסכמה חכמה של תלויויות

AI מנתח תלויויות בחבילות ומצא:
- תלויויות חסרות
- ניגודים בגרסאות
- חבילות מיותרות

```bash
# בדוק תלויויות עם AI
apx ai-check-deps subsystem-name

# קבל הצעות להסרת חבילות מיותרות
apx ai-cleanup-deps subsystem-name

# קבל דוח AI על בריאות המערכת
apx ai-health-report
```

### רצועת זמן נתמכת

AI תומך בעברית, אנגלית ועוד 30+ שפות:

```bash
# שנה שפה לעברית
apx --language=he ai-recommend package-name

# קבל סיוע AI בשפתך
apx ai-help --lang=he
```

## אבטחה והאימות

Apx כוללת מנגנוני אבטחה מקיפים להגנה על המערכת שלך ועל נתוניך.

### ניהול מפתחות API

```bash
# יצירת מפתח API חדש
apx api create-key --name="my-api-key" --permissions="read,write"

# רישום כל מפתחות ה-API שלך
apx api list-keys

# מחיקה של מפתח API
apx api delete-key api-key-id

# סיבוב מפתח API (עדכון לאבטחה)
apx api rotate-key api-key-id
```

### אימות מבוסס-חשבון

```bash
# התחברות למערכת Apx
apx auth login --username=your-username

# התחברות עם אימות דו-שלבי
apx auth login --username=your-username --2fa

# התנתקות
apx auth logout

# בדיקת סטטוס ההתחברות הנוכחי
apx auth status
```

### הנהלת הרשאות

```bash
# הצגת הרשאות המשתמש הנוכחי
apx auth permissions list

# הקצאת הרשאות למשתמש אחר
apx auth permissions grant --user=username --permission=install

# שלילת הרשאות
apx auth permissions revoke --user=username --permission=delete
```

### קידוד מעברים עבור SSH

```bash
# יצירת מפתח SSH חדש
apx ssh-keygen --key-name="secure-key" --key-type=rsa --bits=4096

# שמירת מפתח שדור ציבורי
apx ssh-keygen --public-key --key-name="secure-key"

# הסרת מפתח SSH
apx ssh-keygen --delete --key-name="secure-key"
```

## תמיכה ב-Windows ו-macOS

Apx תומך בפלטפורמות חוצות-מערכות עם שלי עבודה מלאה על Windows ו-macOS.

### Windows Subsystem for Linux (WSL)

```bash
# התקנה ב-WSL2
wsl --install -d Ubuntu
wsl
cd Vanilla-
make build
sudo make install

# הרצה של Apx דרך Windows
wsl apx subsystems list

# גישה לקבצים משותפים
apx ubuntu-latest install /mnt/c/Users/YourName/projects
```

### macOS - Homebrew

```bash
# התקנה דרך Homebrew
brew tap vanilla-os/apx
brew install apx

# עדכון ל-Homebrew
brew upgrade apx

# הסרה
brew uninstall apx
```

### Docker על Windows/macOS

```bash
# התקנת Docker Desktop
# https://www.docker.com/products/docker-desktop

# בנייה של Apx בתוך קונטיינר Docker
docker build -t apx:latest .

# הרצת Apx
docker run -it apx:latest apx subsystems list

# שמירת תמונה לשימוש עתידי
docker save apx:latest > apx-image.tar
```

### שיתוף פקודות בין פלטפורמות

```bash
# פקודות זהות ב-Windows, macOS, ו-Linux
apx subsystems list
apx subsystems new --name=ubuntu-latest
apx ubuntu-latest install python3

# סנכרון קבצים חוצה-פלטפורמה
apx ubuntu-latest sync /local/path /remote/path
```

## CI/CD Workflows - GitHub Actions

Apx משתלב עם GitHub Actions לאוטומציה מלאה של בדיקות והפצות.

### סרטון GitHub Actions בסיסי

```yaml
# .github/workflows/ci.yml
name: CI Pipeline

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.25.0'
    
    - name: Build with Apx
      run: |
        make build
    
    - name: Run Tests
      run: |
        make test
    
    - name: Run Linter
      run: |
        make lint
```

### סרטון Release Automation

```yaml
# .github/workflows/release.yml
name: Release

on:
  push:
    tags:
      - 'v*'

jobs:
  release:
    runs-on: ubuntu-latest
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Build Release
      run: |
        make release
    
    - name: Create Release
      uses: actions/create-release@v1
      with:
        tag_name: ${{ github.ref }}
        release_name: Release ${{ github.ref }}
        draft: false
        prerelease: false
```

### סרטון בדיקות יומיות

```yaml
# .github/workflows/nightly-tests.yml
name: Nightly Tests

on:
  schedule:
    - cron: '0 2 * * *'

jobs:
  nightly:
    runs-on: ubuntu-latest
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Run Comprehensive Tests
      run: |
        make test-all
    
    - name: Generate Coverage Report
      run: |
        make coverage
    
    - name: Upload to Codecov
      uses: codecov/codecov-action@v3
```

## תיעוד API

Apx מספקת API ממוקדת לתכנות המערכת באופן קולחותו.

### נקודות קצה REST בסיסיות

```bash
# רישום כל התת-מערכות
curl -H "Authorization: Bearer YOUR_API_KEY" \
  https://api.apx.local/subsystems

# יצירת תת-מערכת חדשה
curl -X POST -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"name":"ubuntu-latest"}' \
  https://api.apx.local/subsystems

# התקנת חבילה בתת-מערכת
curl -X POST -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"package":"python3"}' \
  https://api.apx.local/subsystems/ubuntu-latest/install

# מחיקת תת-מערכת
curl -X DELETE -H "Authorization: Bearer YOUR_API_KEY" \
  https://api.apx.local/subsystems/ubuntu-latest
```

### דוגמאות JSON

```json
// רישום תת-מערכות
GET /api/subsystems
Response:
{
  "subsystems": [
    {
      "id": "ubuntu-latest",
      "name": "ubuntu-latest",
      "status": "running",
      "packages": 42
    }
  ]
}

// יצירת תת-מערכת
POST /api/subsystems
Body:
{
  "name": "fedora-37",
  "stack": "fedora"
}
Response:
{
  "success": true,
  "subsystem_id": "fedora-37"
}
```

### GraphQL API

```graphql
# שאילתה - רישום תת-מערכות
query {
  subsystems {
    id
    name
    status
    packages {
      name
      version
    }
  }
}

# מוטציה - יצירת תת-מערכת
mutation {
  createSubsystem(input: {
    name: "alpine-latest"
    stack: "alpine"
  }) {
    success
    subsystem {
      id
      name
    }
  }
}
```

## קבצי תצורה

### config.yaml - קבוצה ראשית

```yaml
# ~/.config/apx/config.yaml
version: 1

# הגדרות כלליות
general:
  default_stack: ubuntu
  auto_update: true
  check_updates_interval: 604800  # שבועי

# הגדרות אבטחה
security:
  require_auth: true
  enable_2fa: false
  ssh_key_location: ~/.ssh/apx_key
  api_key_expiry: 2592000  # 30 ימים

# הגדרות רשתות
network:
  enable_remote_access: true
  remote_port: 8080
  use_ssl: true
  ssl_cert_path: /etc/apx/certs/cert.pem
  ssl_key_path: /etc/apx/certs/key.pem

# הגדרות AI
ai:
  enabled: true
  model: gpt-3.5-turbo
  api_key_env: APX_AI_KEY
  language: he

# הגדרות רישום
logging:
  level: info
  file: ~/.apx/apx.log
  max_size: 10485760  # 10MB
  max_backups: 5
```

### .env - משתנים סביבתיים

```bash
# ~/.apx/.env
# מפתחות API
APX_API_KEY=your-secret-api-key
APX_AI_API_KEY=your-ai-api-key

# מידע מארח מרוחק
APX_REMOTE_HOST=example.com
APX_REMOTE_USER=username
APX_REMOTE_PORT=22

# הגדרות רישום
APX_LOG_LEVEL=debug
APX_LOG_FILE=/var/log/apx/apx.log

# הגדרות SSL
APX_SSL_CERT=/etc/apx/certs/cert.pem
APX_SSL_KEY=/etc/apx/certs/key.pem
```

### docker-compose.yml - פריסה

```yaml
version: '3.8'

services:
  apx-server:
    image: apx:latest
    container_name: apx-server
    ports:
      - "8080:8080"
    environment:
      - APX_API_KEY=${APX_API_KEY}
      - APX_AI_API_KEY=${APX_AI_API_KEY}
      - APX_LOG_LEVEL=info
    volumes:
      - apx-data:/var/lib/apx
      - apx-config:/etc/apx
    restart: unless-stopped

  apx-redis:
    image: redis:latest
    container_name: apx-redis
    ports:
      - "6379:6379"
    volumes:
      - redis-data:/data
    restart: unless-stopped

volumes:
  apx-data:
  apx-config:
  redis-data:
```

## תלויויות

### תלויויות ישירות
- `github.com/google/uuid` - יצירת UUID
- `github.com/vanilla-os/sdk` - Framework יישום
- `gopkg.in/yaml.v2` - תצורת YAML

### תלויויות ממשק משתמש
- `github.com/charmbracelet/bubbletea` - Framework ממשק טרמינל
- `github.com/charmbracelet/lipgloss` - עיצוב טרמינל

### תלויויות AI
- `github.com/tmc/langchain-go` - שיתוף פעולה LLM
- `github.com/sashabaranov/go-openai` - שיתוף פעולה OpenAI (אופציונלי)
- `github.com/google/generative-ai-python` - שיתוף פעולה Google AI (אופציונלי)

### תלויויות אבטחה
- `github.com/golang-jwt/jwt/v5` - JWT tokens
- `golang.org/x/crypto` - קריפטוגרפיה

## תיעוד

לתיעוד מפורט, בקר ב: https://docs.vanillaos.org/docs/he/apx

## תרגומים

תרום תרגומים דרך [Weblate](https://hosted.weblate.org/projects/vanilla-os/apx)

## רישיון

GNU General Public License v3.0

## מחברים

- Mirko Brombin <brombin94@gmail.com>
- Pietro di Caprio <pietro@fabricators.ltd>
- תורמי Vanilla OS

## תרומה

תרומות מתקבלות בברכה! אנא ודא:

1. תלויויות מעודכנות עם `go get`, `go mod tidy`, ו-`go mod vendor`
2. הקוד עוקב אחר קונבנציות Go
3. תרגומים מעודכנים ב-Weblate
4. בדיקות עוברות בהצלחה עם `make test`
5. קוד עוקב אחר הסטנדרטים שנבדקו עם `make lint`
